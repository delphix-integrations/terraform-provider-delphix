# docs

This folder contains user-facing documentation published to the [Terraform Registry](https://registry.terraform.io/providers/delphix-integrations/delphix/latest/docs).

## Structure

```
docs/
├── index.md              # Provider overview: auth, system requirements, example usage
├── guides/
│   └── provider_guide.md # Links to examples and community resources
└── resources/
    ├── vdb.md
    ├── vdb_group.md
    ├── environment.md
    ├── appdata_dsource.md
    ├── oracle_dsource.md
    ├── database_postgresql.md
    ├── database_plugin.md
    ├── engine_configuration.md
    └── engine_dct_registration.md
```

## Writing or Updating a Resource Doc

Use `/generate-doc <resource-name>` to scaffold or refresh a doc. Each `docs/resources/<name>.md` must follow this section order:

1. **Title** — `# Resource: delphix_<name>`
2. **Description** — one paragraph explaining what the resource manages
3. **Notes** — bullet list of caveats (computed fields, unsupported updates, version restrictions)
4. **Example Usage** — minimal working HCL with `UPPER_SNAKE_CASE` placeholders; include a `timeouts` block if the resource supports it
5. **Argument Reference** — all schema fields grouped logically; mark updatable fields with `[Updatable]`; document nested blocks as sub-sections
6. **Timeout Configuration** — standard create/update/delete block (default 20m each); explain behavior on timeout per operation
7. **Import** — `import {}` block example using the resource ID; omit if import is not supported
8. **Limitations** — non-updatable fields, unsupported operations, version-gated behaviour

## Conventions

- HCL placeholders use `UPPER_SNAKE_CASE` (e.g. `SOURCE_DATA_ID`, `DATASET_GROUP_ID`).
- Mark every field that appears in `updatable<Name>Keys` in `commons.go` with `[Updatable]`.
- Fields that trigger recreation (`isDestructive<Name>Update = true`) should be noted as non-updatable.
- Version-restricted fields (e.g. DCT v2025.1 only) must include an explicit version note.
- Sensitive fields (passwords, keys) should reference Terraform's sensitive input variables docs.

## Updating index.md

Update `docs/index.md` when:
- The minimum supported DCT or Continuous Data Engine version changes (see **System Requirements** table).
- A new provider-level configuration parameter is added to `provider.go`.
- A new resource is added (the registry sidebar is auto-generated, but the example usage block may need updating).
