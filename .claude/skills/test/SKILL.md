---
name: test
description: Run unit tests for the Delphix Terraform provider. Use when the user wants to run tests, check test results, or verify code correctness.
disable-model-invocation: true
---

Run unit tests for the Delphix Terraform provider.

1. Run `make test` to execute all unit tests (parallel=4, timeout=30s)
2. To run a specific test: `go test ./internal/provider -run $ARGUMENTS -v`
3. Report any failures with the full test output
4. If no argument is provided, run all unit tests with `make test`

Note: Acceptance tests (`make testacc`) require `DCT_KEY` and `DCT_HOST` env vars — use the `/run-testacc` skill for those.
