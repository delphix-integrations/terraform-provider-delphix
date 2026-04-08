---
name: build
description: Build and install the Delphix Terraform provider. Use when the user wants to compile, build, or install the provider binary.
disable-model-invocation: true
---

Build and install the Delphix Terraform provider.

1. Run `make build` to compile the provider binary (`./terraform-provider-delphix`)
2. If the user wants to install it locally, run `make install` to place it at `~/.terraform.d/plugins/delphix.com/dct/delphix/4.2.1/darwin_arm64`
3. For multi-platform release binaries, run `make release` (outputs to `./bin/`)

Report any build errors clearly and suggest fixes.
