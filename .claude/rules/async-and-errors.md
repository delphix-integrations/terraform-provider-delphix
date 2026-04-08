# Async Operations and Error Handling Rules

## Async Job Polling

All DCT API operations that return a job ID must be polled until a terminal state is reached.

```go
status, errMsg := PollJobStatus(jobId, ctx, client)
if errMsg != "" {
    return diag.Errorf("job failed: %s", errMsg)
}
```

Terminal states (constants in `commons.go`): `COMPLETED`, `FAILED`, `TIMEDOUT`, `CANCELED`, `ABANDONED`.

Default Create/Update/Delete timeout: **20 minutes**. Always set `Timeouts` in the resource schema.

## API Error Handling

Call `apiErrorResponseHelper` after every DCT SDK call — never ignore the HTTP response:

```go
res, httpRes, err := client.SomeAPI.SomeOperation(ctx).Execute()
if err := apiErrorResponseHelper(ctx, res, httpRes, err); err != nil {
    return err
}
```

## Object Existence Checks

- In Read: use `PollForObjectExistence()` — do not make a bare API call.
- On 404: call `PollForObjectDeletion()` then `d.SetId("")` to remove from state.
- In Delete: call `PollForObjectDeletion()` after issuing the delete to confirm removal.
