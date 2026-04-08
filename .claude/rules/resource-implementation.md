# Resource Implementation Rules

Rules for implementing Terraform resources in this provider.

## Schema

- Every resource function must have `Description`, `CreateContext`, `ReadContext`, `DeleteContext`, and `Timeouts` (20-min defaults).
- Add `UpdateContext` only when at least one field is updatable in-place (not `ForceNew`).
- Always set `CustomizeDiff: CustomizeDiffTags` unless the resource has no `tags` field.
- `id` field must be `Computed: true`.
- Use `Sensitive: true` on all password/credential fields.
- Use `ForceNew: true` on fields that cannot be updated without destroying the resource.
- Use `Computed: true` on fields returned by the API but not supplied by the user.

## Create

- After every API call, check errors with `apiErrorResponseHelper(ctx, apiRes, httpRes, err)`.
- Set the resource ID with `d.SetId(...)` before polling.
- Call `PollJobStatus()` for every operation that returns a job ID.
- End Create by delegating to the Read function — do not repeat `d.Set()` calls.

## Read

- Handle 404 gracefully: call `PollForObjectDeletion()` and set `d.SetId("")` to signal the resource is gone.
- Set all computed/read-back fields with `d.Set()`.
- Do NOT call `d.SetId()` inside Read — the SDK manages the ID.

## Update

- Check `updatable<Name>Keys` map before issuing any API call.
- Only call the API if `d.HasChange(key)` is true for that key.
- End Update by delegating to the Read function.

## Delete

- Call `apiErrorResponseHelper()` on the delete API call.
- Call `PollForObjectDeletion()` to confirm the resource is gone.
- Do NOT call `d.SetId("")` — the SDK sets this automatically after Delete returns nil.

## commons.go

Every new resource must add two maps:
- `updatable<Name>Keys map[string]bool` — keys that can be updated in-place (value `true`) or not at all (value `false`).
- `isDestructive<Name>Update map[string]bool` — keys whose in-place update requires a destroy+recreate cycle (`true`).
Every key present and `true` in `updatable<Name>Keys` must also appear in `isDestructive<Name>Update`.

## provider.go

Register every new resource under `"delphix_<name>"` in `ResourcesMap`.
