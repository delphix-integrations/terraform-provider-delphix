---
name: sync-sdk
description: Upgrade the DCT Go SDK to the latest version and report what changed. Updates go.mod/go.sum, runs go mod tidy, builds the provider, and summarizes new or changed APIs relevant to this provider.
---

Upgrade the DCT SDK dependency for this Terraform provider.

Follow these steps:

1. **Check current version**
   ```bash
   grep "dct-sdk-go" go.mod
   ```

2. **Find the latest available version**
   ```bash
   go list -m -versions github.com/delphix/dct-sdk-go/v25
   ```

3. **Upgrade to the latest version**
   ```bash
   go get github.com/delphix/dct-sdk-go/v25@latest
   go mod tidy
   ```

4. **Build to catch breaking changes**
   ```bash
   make build
   ```
   If the build fails, report each compile error with the affected file and a suggested fix.

5. **Summarize what changed** — diff the SDK model/API files to find:
   - New API methods (new resources we could implement)
   - New fields on existing models (schema additions needed)
   - Removed/renamed methods (breaking changes to fix)

   Use the `sdk-researcher` agent to explore any new resource types found.

6. Report the upgrade result:
   - Old version → New version
   - Build status: PASS / FAIL (with errors)
   - New APIs available (with brief description)
   - Breaking changes found (with fix suggestions)
