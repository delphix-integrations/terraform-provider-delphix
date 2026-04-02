---
name: run-testacc
description: Run a specific acceptance test for the Delphix Terraform provider. Use when the user wants to run TF_ACC acceptance tests against a real DCT instance.
disable-model-invocation: true
argument-hint: "TestFunctionName"
---

Run a specific acceptance test against a live DCT instance.

The test name to run is: $ARGUMENTS

Steps:
1. Verify `DCT_HOST` and `DCT_KEY` environment variables are set. If not, remind the user to set them in `.claude/settings.local.json` or export them in the shell.
2. Run: `TF_ACC=1 go test ./internal/provider -run $ARGUMENTS -v -timeout 120m`
3. Report the full test output, highlighting PASS/FAIL clearly.

Note: Some tests require additional env vars (e.g., `DATASOURCE_ID`, `ENVIRONMENT_ID`). Check the test function for required vars if the test fails with missing config errors.
