# Tag Handling Rules

## Standard Tag Schema

Every resource that supports tags must include:
- A `tags` field using the standard list-of-key-value block pattern.
- An `ignore_tag_changes` bool field to let users suppress tag drift detection.
- `make_current_account_owner` bool field where applicable.

## CustomizeDiff

Apply `CustomizeDiffTags` to every resource that has a `tags` field:

```go
func resourceFoo() *schema.Resource {
    return &schema.Resource{
        // ...
        CustomizeDiff: CustomizeDiffTags,
    }
}
```

Do not skip `CustomizeDiff` on resources with tags — without it, tag drift will not be detected correctly when `ignore_tag_changes` is set.
