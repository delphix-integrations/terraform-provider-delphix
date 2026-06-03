# docs/

Auto-generated Terraform Registry documentation. **Do not edit files in this directory by hand** — they are generated from schema `Description` fields in [internal/provider/](../internal/provider/).

## Contents

| Path | Purpose |
|---|---|
| [index.md](index.md) | Provider-level docs: authentication, required env vars, `DCT_KEY`/`DCT_HOST` configuration |
| [resources/](resources/) | One `.md` per resource, generated from schema |
| [guides/](guides/) | Hand-written usage guides published to the Registry |

## Regenerating Docs

Docs are auto-generated. To regenerate after changing a resource schema:
```bash
go generate ./...
```
or run `tfplugindocs generate` if the tool is installed separately.

## Rules

- Keep `Description` fields in resource schemas accurate — they appear verbatim in the Registry docs.
- Do not add new `.md` files here manually unless adding a new guide under `guides/`.
- Guides in `guides/` are hand-authored and should link to the corresponding `examples/` directory.
