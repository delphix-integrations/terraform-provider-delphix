## Contributing
 1. Fork the project.
 1. Make your bug fix or new feature.
 1. Add tests for your code.
 1. Send a pull request.

Contributions must be signed as `User Name <user@email.com>`. Make sure to [set up Git with user name and email address](https://git-scm.com/book/en/v2/Getting-Started-First-Time-Git-Setup). Bug fixes should branch from the current stable branch. New features should be based on the release branch.

## Contributor Agreement
All contributors are required to sign the Delphix Contributor agreement prior to contributing code to an open source repository. This process is handled automatically by [cla-assistant](https://cla-assistant.io/). Simply open a pull request and a bot will automatically check to see if you have signed the latest agreement. If not, you will be prompted to do so as part of the pull request process.

## Code of Conduct
This project operates under the Delphix Code of Conduct. By participating in this project you agree to abide by its terms.

## CI

All pull requests to `main` or `develop` must pass the `ci / unit-tests` GitHub Actions check before merging.
The check runs automatically when you open or update a PR.

### What the CI Check Does

1. Checks out your branch at the PR commit SHA.
2. Installs Go (version auto-detected from `go.mod`).
3. Runs `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s`.
4. Uploads `coverage.out` as a downloadable artifact (7-day retention).
5. Fails if total coverage falls below the configured threshold
   (see `COVERAGE_THRESHOLD` in `.github/workflows/ci.yml`).

### Verifying Locally Before Pushing

```bash
go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s
go tool cover -func=coverage.out | tail -1
```

### Acceptance Tests

Acceptance tests (`TestAcc*`) require a live DCT instance and are **not run in CI** —
they are excluded automatically because CI does not set `TF_ACC=1`.
See `## Build & Test Commands` in `CLAUDE.md` for how to run acceptance tests locally.
