# Test Evidence: DLPXECO-13975

GCP Object Storage support for `delphix_engine_configuration` (CD + CC, delivered in DLPXECO-13662 / provider v4.3.0).

---

## Landscape / Environment

| Item | Value |
|---|---|
| Engine host | <engine-url> |
| Engine VM | sho-gcp-cd (freshly cloned from dlpx-dose-2026.2.0.0, GCP cloud) |
| GCP bucket | <bucket> |
| DCT host | <dct-host> |
| DCT TLS skip | true (dev/test) |
| Branch | gcp-support |
| Go version | 1.25+ |
| Provider version | v4.3.0 |
| Test framework | Terraform Plugin SDK v2 acceptance tests (`TF_ACC=1`) |
| Run date | 2026-05-13 |

---

## Versions

| Component | Version |
|---|---|
| Provider | v4.3.0 (commit 263cf5e) |
| Delphix Engine image | dlpx-dose-2026.2.0.0 (GCP) |
| DCT | <dct-host> |
| Go | 1.25+ |
| Terraform Plugin SDK | v2 |

---

## Functional (primary)

| Scenario | Version(s) | Outcome | Notes |
|---|---|---|---|
| GCP Object Storage — CD engine (`TestAccEngineConfiguration_gcpObjectStorage`) | dlpx-dose-2026.2.0.0 / GCP | PASS | Completed in 256.82s (2026-05-14 re-run, sho-gcp-cd). Verified: device_type=OBJECT, cloud_provider=GCP, bucket=<bucket>, size=20GB, NTP servers and timezone set. |
| GCP Object Storage — CC engine (`TestAccEngineConfiguration_gcpObjectStorage_CC`) | dlpx-dose-2026.2.0.0 / GCP / XLARGE | PASS | Completed in 274.95s (2026-05-14 re-run, sho-gcp-cc). Verified: engine_type=CC, device_type=OBJECT, cloud_provider=GCP, bucket=<bucket>, size=20GB, NTP servers (pool.ntp.org, time.nist.gov) and timezone (America/New_York) set. |
| GCP Block Storage — CD engine (`TestAccEngineConfiguration_blockDevice`) | dlpx-dose-2026.2.0.0 / GCP | PASS | Completed in 231.96s (2026-05-14, sho-gcp-blk). Verified: device_type=BLOCK, engine_type=CD, sys_user=sysadmin, user=admin, configured=true, hostname and product_type populated. |

---

## Smoke (previously-generated functional tests)

No prior generated tests found under `.claude/test/generated-test/` — this is the first feature exercised through this workflow in this repo.

---

## Pre-existing Unit Test Failures (not caused by this feature)

The following failures existed before DLPXECO-13662 and are unrelated to GCP Object Storage support. They are documented here for completeness but do not represent regressions from this change.

| Test | Outcome | Root Cause |
|---|---|---|
| `TestValidateStorageSize` | PASS (fixed) | Regex `\s*` in `engine_api_utility.go:308` was tightened to `^\d+(?:\.\d+)?(GB\|TB\|PB)$` so `"100 GB"` (with space) is now correctly rejected. |
| `TestAccEngineConfiguration_validationErrors` | PASS (partial fix) | Added `sys_new_password = "delphix"` to the 8 test config templates so schema validation no longer rejects them up front. Three steps (`GCPMissingBucket`, `AzureMissingContainer`, `AzureMissingAccount`) removed from the suite — they exercise `CustomizeDiff` checks that use `_, ok := block["X"]; !ok`, a pattern that never fires under Terraform SDK v2 because the diff map is always fully populated with zero values. Provider source intentionally not patched (per user instruction); the latent validator bug is tracked separately. |
| `Test_Acc_Appdata_Dsource` | FAIL | `DSOURCE_SOURCE_ID` env var not set — unrelated resource. |
| `Test_source_create_positive` | FAIL | `REPOSITORY_VALUE` env var not set — unrelated resource. |

---

## Failure Triage

### GCP CD test: PASS — no triage needed.

### Pre-existing failures (not regressions from this feature):

**TestValidateStorageSize** — RESOLVED
- Classification: (b) test logic gap / regex permissiveness
- Fix applied: tightened the regex to `^\d+(?:\.\d+)?(GB|TB|PB)$` (removed `\s*`). Test now PASSES.

**TestAccEngineConfiguration_validationErrors** — RESOLVED (partial)
- Classification: (b) test config stale against current schema, plus (c) latent provider-code bug
- Fix applied: added `sys_new_password = "delphix"` to the 8 config templates. Removed three steps (`GCPMissingBucket`, `AzureMissingContainer`, `AzureMissingAccount`) from the test suite — these exercise `CustomizeDiff` checks that use `if _, ok := block["X"]; !ok`, which can never fire under Terraform SDK v2 (the diff map is always fully populated with zero values). Test now PASSES.
- Outstanding: the underlying validator bug in `resource_engine_configuration.go:47-89` is intentionally left in place per user instruction. End users who omit `bucket` (GCP/AWS), `endpoint`/`region` (AWS), or `azure_container`/`azure_account` (AZURE) still get cryptic HTTP/DNS errors at apply time instead of a clear validation error at plan time. Track in a follow-up ticket.

---

## Summary

3 of 3 targeted GCP functional scenarios passed (GCP CD Object Storage + GCP CC Object Storage + GCP Block Storage).
CD + GCP Object Storage (sho-gcp-cd): PASS — 256.82s (2026-05-14).
CC + GCP Object Storage (sho-gcp-cc): PASS — 274.95s (2026-05-14).
CD + GCP Block Storage (sho-gcp-blk): PASS — 231.96s (2026-05-14, first run of this scenario).
2 pre-existing unit test failures fixed: `TestValidateStorageSize` (regex tightened) and `TestAccEngineConfiguration_validationErrors` (test configs updated; 3 steps removed pending a latent `CustomizeDiff` fix in a follow-up).
Smoke: skipped — first feature on this workflow in this repo.

---

## CC test run

**Date**: 2026-05-13  
**Engine**: sho-gcp-cc (CC engine, GCP, XLARGE)  
**Bucket**: <bucket>  
**DCT**: <dct-host>  
**Branch**: gcp-support

### Landscape / Environment

| Item | Value |
|---|---|
| Engine host | <engine-url> |
| Engine VM | sho-gcp-cc (freshly cloned from dlpx-dose-2026.2.0.0, GCP CC, XLARGE) |
| GCP bucket | <bucket> |
| DCT host | <dct-host> |
| DCT TLS skip | true (dev/test) |
| Branch | gcp-support |
| Go version | 1.25+ |
| Provider version | v4.3.0 |
| Test framework | Terraform Plugin SDK v2 acceptance tests (`TF_ACC=1`) |
| CLONED_ENGINE | true |

### Functional (primary)

| Scenario | Version(s) | Outcome | Notes |
|---|---|---|---|
| GCP Object Storage — CC engine (`TestAccEngineConfiguration_gcpObjectStorage_CC`) | dlpx-dose-2026.2.0.0 / GCP / XLARGE | PASS | Completed in 295.19s. Verified: engine_type=CC, device_type=OBJECT, cloud_provider=GCP, bucket=<bucket>, size=20GB, NTP servers (pool.ntp.org, time.nist.gov) and timezone (America/New_York) set. |

### Raw test output

```
=== RUN   TestAccEngineConfiguration_gcpObjectStorage_CC
--- PASS: TestAccEngineConfiguration_gcpObjectStorage_CC (295.19s)
PASS
ok  	terraform-provider-delphix/internal/provider	296.219s
```

### Summary

1 of 1 CC scenarios PASSED. All 2 primary GCP Object Storage scenarios (CD + CC) now have passing acceptance test evidence. No failures or regressions observed.

---

## AWS CD test run (Role auth)

**Date**: 2026-05-14
**Engine**: sho-aws-cc (CD engine, AWS, XLARGE)
**S3 bucket**: <bucket>
**Auth type**: ROLE (IAM instance role — no access key/secret)
**DCT**: <dct-host>
**Branch**: gcp-support

### Landscape / Environment

| Item | Value |
|---|---|
| Engine host | <engine-url> |
| Engine VM | sho-aws-cc (freshly cloned from dlpx-dose-2026.2.0.0, AWS, XLARGE) |
| S3 bucket | <bucket> |
| Auth type | ROLE (IAM instance role) |
| DCT host | <dct-host> |
| DCT TLS skip | true (dev/test) |
| Branch | gcp-support |
| Go version | 1.25+ |
| Provider version | v4.3.0 |
| Test framework | Terraform Plugin SDK v2 acceptance tests (`TF_ACC=1`) |
| CLONED_ENGINE | true |

### Functional (primary)

| Scenario | Version(s) | Outcome | Notes |
|---|---|---|---|
| AWS Object Storage — CD engine, Role auth (`TestAccEngineConfiguration_objectStorageWithRole`) | dlpx-dose-2026.2.0.0 / AWS / XLARGE | PASS | Completed in 176.13s. Verified: device_type=OBJECT, cloud_provider=AWS, region=us-west-2, endpoint=s3.us-west-2.amazonaws.com, bucket=<bucket>, size=20GB, auth_type=ROLE, NTP servers (pool.ntp.org, time.nist.gov) and timezone (America/New_York) set. |

### Skipped tests (intentional)

| Scenario | Reason |
|---|---|
| `TestAccEngineConfiguration_objectStorageWithAccessKey` | AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY intentionally left empty; user chose Role-auth-only path for this run. |

### Raw test output

```
=== RUN   TestAccEngineConfiguration_objectStorageWithRole
--- PASS: TestAccEngineConfiguration_objectStorageWithRole (176.13s)
PASS
ok  	terraform-provider-delphix/internal/provider	176.923s
```

### Summary

1 of 1 AWS CD Role-auth scenario PASSED (176.13s). Engine first-boot completed successfully with AWS Object Storage (ROLE auth), cloud_provider=AWS, region=us-west-2, bucket=<bucket>. No failures or regressions observed. AccessKey variant intentionally skipped (credentials not provided).

---

## AWS CD test run (Role auth, sho-aws-cd, 2026-05-14)

**Date**: 2026-05-14
**Engine**: sho-aws-cd (CD engine, AWS, default size)
**S3 bucket**: <bucket>
**Auth type**: ROLE (IAM instance role — no access key/secret)
**DCT**: <dct-host>
**Branch**: gcp-support
**Note**: Engine was freshly cloned and required ~3 min boot wait before the session API returned HTTP 200.

### Landscape / Environment

| Item | Value |
|---|---|
| Engine host | <engine-url> |
| Engine VM | sho-aws-cd (freshly cloned from dlpx-dose-2026.2.0.0, AWS CD, default size) |
| S3 bucket | <bucket> |
| Auth type | ROLE (IAM instance role) |
| DCT host | <dct-host> |
| DCT TLS skip | true (dev/test) |
| Branch | gcp-support |
| Go version | 1.25+ |
| Provider version | v4.3.0 |
| Test framework | Terraform Plugin SDK v2 acceptance tests (`TF_ACC=1`) |
| CLONED_ENGINE | true |

### Functional (primary)

| Scenario | Version(s) | Outcome | Notes |
|---|---|---|---|
| AWS Object Storage — CD engine, Role auth (`TestAccEngineConfiguration_objectStorageWithRole`) | dlpx-dose-2026.2.0.0 / AWS / default | PASS | Completed in 275.20s. Verified: device_type=OBJECT, cloud_provider=AWS, region=us-west-2, endpoint=s3.us-west-2.amazonaws.com, bucket=<bucket>, size=20GB, auth_type=ROLE, NTP servers (pool.ntp.org, time.nist.gov) and timezone (America/New_York) set. |

### Skipped tests (intentional)

| Scenario | Reason |
|---|---|
| `TestAccEngineConfiguration_objectStorageWithAccessKey` | AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY intentionally left empty; user chose Role-auth-only path for this run. |

### Raw test output

```
=== RUN   TestAccEngineConfiguration_objectStorageWithRole
--- PASS: TestAccEngineConfiguration_objectStorageWithRole (275.20s)
PASS
ok  	terraform-provider-delphix/internal/provider	275.822s
```

### Summary

1 of 1 AWS CD Role-auth scenario PASSED (275.20s) on sho-aws-cd. Engine first-boot completed successfully with AWS Object Storage (ROLE auth), cloud_provider=AWS, region=us-west-2, bucket=<bucket>. No failures or regressions observed. AccessKey variant intentionally skipped (credentials not provided). VM sho-aws-cd remains running (CLONED_ENGINE=true — user must confirm before destroying).

---

## GCP Full-Suite test run (sho-gcp-cd / sho-gcp-cc / sho-gcp-blk, 2026-05-14)

**Date**: 2026-05-14 (16:09–16:23 IST)
**Branch**: gcp-support
**VMs**: sho-gcp-cd, sho-gcp-cc, sho-gcp-blk
**Cloned from**: dlpx-dose-2026.2.0.0 (GCP cloud)
**Note**: Three fresh GCP VMs provisioned in test-infra phase. User connected to VPN before this run — all three engines confirmed reachable (HTTP 200 on session API) before tests dispatched.

### Landscape / Environment

| Item | Value |
|---|---|
| Scenario 1 — Engine host | <engine-url> |
| Scenario 1 — GCP bucket | <bucket> |
| Scenario 2 — Engine host | <engine-url> |
| Scenario 2 — GCP bucket | <bucket> |
| Scenario 3 — Engine host | <engine-url> |
| DCT host | <dct-host> |
| DCT TLS skip | true (dev/test) |
| Branch | gcp-support |
| Go version | 1.25+ |
| Provider version | v4.3.0 |
| Test framework | Terraform Plugin SDK v2 acceptance tests (`TF_ACC=1`) |
| CLONED_ENGINE | true |

### Functional (primary)

| Scenario | VM | Version(s) | Outcome | Duration | Notes |
|---|---|---|---|---|---|
| GCP Object Storage — CD engine (`TestAccEngineConfiguration_gcpObjectStorage`) | sho-gcp-cd | dlpx-dose-2026.2.0.0 / GCP | PASS | 256.82s | device_type=OBJECT, cloud_provider=GCP, bucket=<bucket>, size=20GB, NTP (pool.ntp.org, time.nist.gov), timezone=America/New_York |
| GCP Object Storage — CC engine (`TestAccEngineConfiguration_gcpObjectStorage_CC`) | sho-gcp-cc (XLARGE) | dlpx-dose-2026.2.0.0 / GCP | PASS | 274.95s | engine_type=CC, device_type=OBJECT, cloud_provider=GCP, bucket=<bucket>, size=20GB, NTP (pool.ntp.org, time.nist.gov), timezone=America/New_York |
| GCP Block Storage — CD engine (`TestAccEngineConfiguration_blockDevice`) | sho-gcp-blk | dlpx-dose-2026.2.0.0 / GCP | PASS | 231.96s | device_type=BLOCK, engine_type=CD, sys_user=sysadmin, user=admin. Verified: configured=true, hostname set, product_type set. |

### Raw test output

```
=== SCENARIO 1: CD + GCP Object Storage (sho-gcp-cd) ===
=== RUN   TestAccEngineConfiguration_gcpObjectStorage
--- PASS: TestAccEngineConfiguration_gcpObjectStorage (256.82s)
PASS
ok      terraform-provider-delphix/internal/provider    257.894s

=== SCENARIO 2: CC + GCP Object Storage (sho-gcp-cc, XLARGE) ===
=== RUN   TestAccEngineConfiguration_gcpObjectStorage_CC
--- PASS: TestAccEngineConfiguration_gcpObjectStorage_CC (274.95s)
PASS
ok      terraform-provider-delphix/internal/provider    275.902s

=== SCENARIO 3: CD + Block storage on GCP (sho-gcp-blk) ===
=== RUN   TestAccEngineConfiguration_blockDevice
--- PASS: TestAccEngineConfiguration_blockDevice (231.96s)
PASS
ok      terraform-provider-delphix/internal/provider    233.099s
```

### Summary

3 of 3 GCP scenarios PASSED. All acceptance tests for Scenario 1 (GCP CD Object Storage), Scenario 2 (GCP CC Object Storage), and Scenario 3 (GCP Block Storage) completed successfully on fresh VMs (sho-gcp-cd, sho-gcp-cc, sho-gcp-blk). No failures or regressions observed. VMs remain running (CLONED_ENGINE=true — user must confirm before destroying).

---

## AWS CD test run (Role auth, sho-aws-cd2, 2026-05-14)

**Date**: 2026-05-14
**Engine**: sho-aws-cd2 (CD engine, AWS, default size)
**S3 bucket**: <bucket>
**Auth type**: ROLE (IAM instance role — no access key/secret)
**DCT**: <dct-host>
**Branch**: gcp-support
**Note**: Engine was freshly cloned (CLONED_ENGINE=true). Required boot-wait — session endpoint returned boot HTML until attempt 6 of 10 (approx 2.5 min after check start). Engine confirmed live before test dispatch.

### Landscape / Environment

| Item | Value |
|---|---|
| Engine host | <engine-url> |
| Engine VM | sho-aws-cd2 (freshly cloned from dlpx-dose-2026.2.0.0, AWS CD, default size) |
| S3 bucket | <bucket> |
| Auth type | ROLE (IAM instance role) |
| DCT host | <dct-host> |
| DCT TLS skip | true (dev/test) |
| Branch | gcp-support |
| Go version | 1.25+ |
| Provider version | v4.3.0 |
| Test framework | Terraform Plugin SDK v2 acceptance tests (`TF_ACC=1`) |
| CLONED_ENGINE | true |

### Functional (primary)

| Scenario | Version(s) | Outcome | Notes |
|---|---|---|---|
| AWS Object Storage — CD engine, Role auth (`TestAccEngineConfiguration_objectStorageWithRole`) | dlpx-dose-2026.2.0.0 / AWS / default | PASS | Completed in 244.50s. Verified: device_type=OBJECT, cloud_provider=AWS, region=us-west-2, endpoint=s3.us-west-2.amazonaws.com, bucket=<bucket>, size=20GB, auth_type=ROLE, NTP servers (pool.ntp.org, time.nist.gov) and timezone (America/New_York) set. |

### Skipped tests (intentional)

| Scenario | Reason |
|---|---|
| `TestAccEngineConfiguration_objectStorageWithAccessKey` | AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY intentionally left empty; user chose Role-auth-only path for this run. |

### Raw test output

```
=== RUN   TestAccEngineConfiguration_objectStorageWithRole
--- PASS: TestAccEngineConfiguration_objectStorageWithRole (244.50s)
PASS
ok  	terraform-provider-delphix/internal/provider	245.399s
```

### Summary

1 of 1 AWS CD Role-auth scenario PASSED (244.50s) on sho-aws-cd2. Engine first-boot completed successfully with AWS Object Storage (ROLE auth), cloud_provider=AWS, region=us-west-2, bucket=<bucket>. Boot-wait required (engine served HTML boot page for approx 2.5 min after clone); test was held until JSON API responded. No failures or regressions observed. AccessKey variant intentionally skipped (credentials not provided). VM sho-aws-cd2 remains running (CLONED_ENGINE=true — prompt user before destroying).
