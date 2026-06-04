# Test Infrastructure — terraform-provider-delphix

This file describes the *infrastructure* required to run `engine_configuration` acceptance tests — how to SSH into the `dc` host that owns the lifecycle of test engines, and how to clone / configure / destroy those engines.

> ## ⚠️ Secret-handling rule (applies to ALL commands in this file)
>
> **Never print, echo, log, or display `DCOA_PASSWORD` (or any other secret env var) in shell output, command text, tool calls, transcripts, or chat messages.**
>
> - Do **not** run commands like `echo $DCOA_PASSWORD`, `env | grep DCOA`, `set`, `printenv DCOA_PASSWORD`, or any command that prints the value.
> - Do **not** interpolate `$DCOA_PASSWORD` into a command string that gets echoed back (e.g. avoid `sshpass -p "$DCOA_PASSWORD" ...` in a way that appears in logs). `DCOA_PASSWORD` is **not** used for the SSH connection itself (SSH uses key auth — see below); it is only supplied to the `dc` application login *inside* the session.
> - Do **not** paste the value into chat — refer to it only by variable name (`$DCOA_PASSWORD`).
> - Any secret a script consumes must be read from `os.Getenv` / equivalent at runtime — never from a command-line argument that shows up in shell history or process listings.
> - The same rule applies to `DCT_KEY`, `DCOA_USER` credentials, the SSH private key (`DCOA_KEY_PATH`), the 2FA TOTP code, and any other sensitive value sourced from `settings.local.json`.

## Provisioning fresh test infrastructure (engine_configuration)

> `engine_configuration` tests need a freshly-cloned **engine** VM — `engine_configuration` is a destructive first-boot config and cannot be re-run against an already-configured engine.

Clone VMs from golden images on the dc host.

### Prerequisite — SSH to the `dc` host

> **NEVER install the `dc` CLI — anywhere, under any circumstance.** The `dc` CLI is already installed and configured on `DCOA_HOST`. Do not run any installer (brew, pip, curl-to-shell, package manager, etc.) on the test runner, dev laptop, or remote host. **The only supported workflow is:** run [test/ssh_dcoa.go](../../test/ssh_dcoa.go) to open an SSH session to `DCOA_HOST`, then issue the relevant `dc` commands inside that session. If `dc` appears to be missing or broken on the remote host, surface that to the user — do not attempt any install workaround.

**Before** running any provisioning step below, open an SSH session to `DCOA_HOST` using the credentials from `.claude/settings.local.json`. Use the helper script [test/ssh_dcoa.go](../../test/ssh_dcoa.go) — it reads `DCOA_HOST` and `DCOA_USER` from the environment and authenticates with an **SSH private key** (key-based / passwordless auth — no SSH password). The key is resolved from `DCOA_KEY_PATH` if set, otherwise the first match under `~/.ssh/` in the order `dcoa_ed25519`, `id_ed25519`, `id_rsa`. The key must be **unencrypted** (or backed by an ssh-agent); encrypted keys are rejected. The script opens the session and runs an arbitrary command on the dc host. All `dc` commands must be issued through this script (or an interactive session it opens) — never locally.

Once the SSH session is established, **log in to `dc`** before any other operation:

- Username: `${DCOA_USER}`
- Password: `${DCOA_PASSWORD}` (pulled from env var; do not retype, do not hardcode)
- Authenticator code: ask the user — 2FA TOTP, rotates every 30s

Workflow when using [test/ssh_dcoa.go](../../test/ssh_dcoa.go) to drive the dc login:

1. The script opens the SSH connection using `DCOA_HOST` + `DCOA_USER` + the SSH private key (`DCOA_KEY_PATH` / `~/.ssh` fallback). SSH itself does **not** use `DCOA_PASSWORD`.
2. It triggers the dc login on the remote host.
3. When the **`dc` application** prompts for the password, supply `DCOA_PASSWORD`. (This is the app login, distinct from the SSH connection — the SSH layer is keyless.)
4. When prompted for the **Authenticator code**, the assistant must **pause and ask the user** for the current TOTP — it is a time-based 2FA token that is **not** stored in `settings.local.json` and cannot be derived from any other env var. Do not guess, do not retry with stale codes.


If `DCOA_HOST` or `DCOA_USER` is empty, or no usable SSH key can be found (`DCOA_KEY_PATH` / `~/.ssh` fallback), halt and ask the user before continuing. `DCOA_PASSWORD` is required for the `dc` app login step — if it is empty when a login is needed, halt and ask the user to populate it.

### Required env vars in `.claude/settings.local.json`

| Var | Required | Purpose | Example |
|---|---|---|---|
| `DCOA_HOST` | **Always** | Hostname of the dc host you SSH into to provision VMs | `dlpxdc.co` |
| `DCOA_USER` | **Always** | SSH user on the dc host (also the `dc` app username) | `user@example.com` |
| `DCOA_KEY_PATH` | Optional | Path to the SSH private key for the dc host. If unset, falls back to `~/.ssh/{dcoa_ed25519, id_ed25519, id_rsa}`. Key must be unencrypted or ssh-agent-backed. | `~/.ssh/dcoa_ed25519` |
| `DCOA_PASSWORD` | **For `dc` login** | Password for the **`dc` application** login on the host (sensitive — gitignored). **Not** used for the SSH connection (SSH is key-based). | — |
| `ENGINE_NAME` | **Always** (source for `DELPHIX_ENGINE_HOST` template) | VM name for the engine — cloned fresh per `engine_configuration` run | `tergcpcc` |
| `ENGINE_GOLDEN_IMAGE` | **Always** | Golden image group used to create a fresh engine VM | `dlpx-develop` |
| `DELPHIX_ENGINE_CLOUD` | **Always** | Target cloud for the cloned VM (also narrows the test regex) | `GCP` |

If any required var is empty, the assistant must halt and ask the user to populate `settings.local.json` before provisioning a VM or running `go test`.

### Derived value: `DELPHIX_ENGINE_HOST`

`DELPHIX_ENGINE_HOST` is stored as the template `"http://${ENGINE_NAME}.dlpxdc.co"` so that updating `ENGINE_NAME` alone updates the engine URL. The assistant **must expand `${ENGINE_NAME}` at run time** when reading the env from `settings.local.json` — substitute the current value of `ENGINE_NAME` before passing the value to `go test`.

Resolution rule:
1. Read `ENGINE_NAME` from `.claude/settings.local.json`.
2. Read the raw `DELPHIX_ENGINE_HOST` template (e.g. `http://${ENGINE_NAME}.dlpxdc.co`).
3. Substitute `${ENGINE_NAME}` with the value from step 1.
4. Export the resolved value as `DELPHIX_ENGINE_HOST` before running tests.

If `ENGINE_NAME` is empty, the URL is incomplete — surface that to the user and ask for a name.

To change the engine target between runs: **edit `ENGINE_NAME` only** — do not edit `DELPHIX_ENGINE_HOST` directly.

### Provisioning workflow

`engine_configuration` tests clone an engine VM from a golden image. VM name, golden image, and cloud target all come from `.claude/settings.local.json` — never hardcoded, never auto-generated.

**Project-side gates (must hold before provisioning):**

1. Required vars (`ENGINE_NAME`, `ENGINE_GOLDEN_IMAGE`, `DELPHIX_ENGINE_CLOUD`) are all non-empty.
2. If a VM already exists with name `ENGINE_NAME`, reuse it — do not re-clone, do not destroy.

**Clone + test + teardown shape:**

1. **Clone a VM** if one with `ENGINE_NAME` doesn't already exist. Use `ENGINE_GOLDEN_IMAGE` as the image, `ENGINE_NAME` as the VM name, and deploy in the `DELPHIX_ENGINE_CLOUD` cloud. Wait for the VM to boot before returning.
2. **Run the test** with the appropriate regex resolved from `DELPHIX_ENGINE_TYPE` / `DELPHIX_ENGINE_CLOUD`:
   ```bash
   go test -v -timeout 120m -run "<regex>" ./internal/provider/
   ```
3. **Teardown** — only destroy the VM if this run created it (tracked via `CLONED_ENGINE`). Always ask the user before deletion.

**Rules:**
- **Idempotent:** if a VM with `ENGINE_NAME` already exists, skip the clone — engines are expensive to provision (3–5 min each).
- **Don't destroy what you didn't create:** only destroy VMs that this run cloned (tracked via `CLONED_ENGINE`).
- **Cloud match:** the engine must be cloned in the cloud matching `DELPHIX_ENGINE_CLOUD` so the test's cloud-specific config (S3/Azure/GCP bucket) is reachable.
- **VM size for CC:** when `DELPHIX_ENGINE_TYPE=CC`, clone with the XLARGE size — the default size is too small for the Continuous Compliance engine to boot reliably. CD engines use the default size.



---

## Selecting the engine_configuration variant via `DELPHIX_ENGINE_TYPE` / `DELPHIX_ENGINE_CLOUD`

`.claude/settings.local.json` has two env vars that compose to narrow which `engine_configuration` test runs:

| Var | Purpose | Values |
|---|---|---|
| `DELPHIX_ENGINE_TYPE` | Engine type | `CD`, `CC`, or empty |
| `DELPHIX_ENGINE_CLOUD` | Storage backend | `BLOCK`, `AWS`, `AZURE`, `GCP`, or empty |

The assistant resolves these to a `go test -run` regex per the table below. Empty narrowers mean "all engine_configuration variants". An explicit chat instruction ("run only the GCP CC test") always overrides the env vars.

| TYPE | CLOUD | `-run` regex | Test function |
|---|---|---|---|
| `CD` | `BLOCK` (or empty CLOUD) | `^TestAccEngineConfiguration_blockDevice$` | block-device CD |
| `CD` | `AWS` | `^TestAccEngineConfiguration_objectStorageWith(Role\|AccessKey)$` | AWS role + access-key |
| `CD` | `AZURE` | `^TestAccEngineConfiguration_azureObjectStorage` | Azure managed-id + access-key |
| `CD` | `GCP` | `^TestAccEngineConfiguration_gcpObjectStorage$` | GCP CD |
| `CC` | `GCP` | `^TestAccEngineConfiguration_gcpObjectStorage_CC$` | GCP CC |
| `CC` | `AWS` / `AZURE` | — | not implemented |
| empty | empty | `^TestAccEngineConfiguration_` | all engine_configuration tests |

### Example invocations

```bash
# DELPHIX_ENGINE_TYPE=CC, DELPHIX_ENGINE_CLOUD=GCP
go test -v -timeout 120m -run "^TestAccEngineConfiguration_gcpObjectStorage_CC$" ./internal/provider/

# DELPHIX_ENGINE_TYPE=CD, DELPHIX_ENGINE_CLOUD=AWS
go test -v -timeout 120m -run "^TestAccEngineConfiguration_objectStorageWith(Role|AccessKey)$" ./internal/provider/

# narrowers empty — every engine_configuration variant
go test -v -timeout 120m -run "^TestAccEngineConfiguration_" ./internal/provider/
```

When a chat message specifies the target ("run only the GCP CC test"), that wins over the env vars.

### Decision questions to ask the user

When the user says "run tests" or "test engine_configuration" without specifying scope, ask:

1. **Which cloud variant?** AWS / Azure / GCP / Block.
2. **CD or CC engine?** CD covers block + all object-storage clouds; CC currently only on GCP.
3. **Is the engine already provisioned?** Existing engine = reuse if not yet first-boot configured; otherwise clone a fresh one.
4. **Destructive OK?** `engine_configuration` first-boot config is **non-reversible** without re-imaging the engine.

### Common DCT env vars

> **Source of truth:** All env vars listed in this file should be populated from [`.claude/settings.local.json`](settings.local.json) under the `env` block. The file is gitignored, so values stay local to your machine. Do **not** hand-export credentials in your shell ad-hoc when `settings.local.json` exists — keep one source of truth so every test run sees the same values.


### End-to-end scenarios

For each scenario: **clone a VM** using golden image `$ENGINE_GOLDEN_IMAGE` and engine name `$ENGINE_NAME` in cloud `$DELPHIX_ENGINE_CLOUD`, wait for the VM to boot before continuing.

- **Scenario 1 — CD + GCP**
  - Clone a VM (default size).
  - Run: `DELPHIX_ENGINE_TYPE=CD DELPHIX_ENGINE_CLOUD=GCP go test -v -timeout 120m -run '^TestAccEngineConfiguration_gcpObjectStorage$' ./internal/provider/`
  - After test completion, ask the user before destroying the engine `$ENGINE_NAME`.

- **Scenario 2 — CC + GCP** (CC requires XLARGE size)
  - Clone a VM with XLARGE size.
  - Run: `DELPHIX_ENGINE_TYPE=CC DELPHIX_ENGINE_CLOUD=GCP go test -v -timeout 120m -run '^TestAccEngineConfiguration_gcpObjectStorage_CC$' ./internal/provider/`
  - After test completion, ask the user before destroying the engine `$ENGINE_NAME`.

- **Scenario 3 — CD + Block** (regression on a GCP-hosted VM)
  - Clone a VM in GCP (default size).
  - Run: `DELPHIX_ENGINE_TYPE=CD DELPHIX_ENGINE_CLOUD=BLOCK go test -v -timeout 120m -run '^TestAccEngineConfiguration_blockDevice$' ./internal/provider/`
  - After test completion, ask the user before destroying the engine `$ENGINE_NAME`.
