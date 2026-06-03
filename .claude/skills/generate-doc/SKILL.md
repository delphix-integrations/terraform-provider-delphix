---
name: generate-doc
description: Generate documentation for a git commit or enhancement. Use when the user wants to create or update CHANGELOG, registry docs, examples, or CLAUDE.md files based on recent code changes.
---

Generate documentation for the most recent git commit, or a commit specified by the user.

1. Identify the commit:
   - If the user provided a hash or range, use that. Otherwise use HEAD.
   - Run `git log -1 --format="%H %s" HEAD` and `git diff HEAD~1 HEAD --stat` to get the overview.
   - Run `git diff HEAD~1 HEAD` to read the full diff.

2. Classify the change based on the diff:
   - Change type: New resource / New field / Bug fix / Enhancement / Refactor / Engine API change
   - Resources affected: which `delphix_*` resources changed
   - Breaking or non-breaking: does this change any existing field name, type, or behavior

3. Apply documentation updates — read each file before editing, then update only what the diff actually changed:

   **A. `CHANGELOG.md`** — Add a new entry at the top. Read `GNUmakefile` for the current `VERSION`. Format:
   ```
   ## <version>
   ### New Features / Enhancements / Bug Fixes
   - `delphix_<resource>`: <what changed and why it matters> (<ticket if in commit message>)
   ```
   If `CHANGELOG.md` does not exist, create it.

   **B. `docs/resources/<name>.md`** — Update the Argument Reference for any new or changed fields. Add a new Example Usage section if a meaningful new scenario was added. Add a version note for new fields: `> **Note:** Requires provider vX.Y.Z or later.`

   **C. `docs/index.md`** — If a new resource or major capability was added, add a row to the Support Matrix table.

   **D. `examples/<resource>/main.tf`** — If the commit adds a new capability not yet shown in the example, add a commented block demonstrating it. Use placeholder credentials (`"1.XXXX"`, `"HOSTNAME"`).

   **E. `CLAUDE.md` files** — Check root `CLAUDE.md`, `examples/<resource>/CLAUDE.md`, and `internal/provider/CLAUDE.md`. Update any section that is now stale.

4. Do NOT edit any `.go` source files.

5. End with a summary table:

   | Artifact | File | Status |
   |---|---|---|
   | CHANGELOG | `CHANGELOG.md` | ✅ Updated / 📄 Created |
   | Registry doc | `docs/resources/<name>.md` | ✅ Updated / ⚠️ No change needed |
   | Index page | `docs/index.md` | ✅ Updated / ⚠️ No change needed |
   | Example config | `examples/<name>/main.tf` | ✅ Updated / ⚠️ Already current |
   | CLAUDE.md files | Various | ✅ Updated / ⚠️ No change needed |

Rules:
- Only document what the diff actually shows — never invent behavior
- Always read a file before editing it
- Mark breaking changes explicitly in both CHANGELOG and registry docs
- Keep examples minimal — show the new feature, not every possible field
- Always state the minimum provider version for new fields
