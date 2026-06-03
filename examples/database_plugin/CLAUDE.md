# examples/database_plugin/

Example for the `delphix_database_plugin` resource — uploading a database plugin (toolkit) to a Delphix Engine via DCT.

## Files

| File | Purpose |
|---|---|
| [main.tf](main.tf) | Upload a plugin artifact to one or more engines |

## When to Use

Database plugins extend Delphix Continuous Data to support additional database types (e.g., MongoDB, custom AppData sources). Use this resource to automate plugin deployment as part of an engine provisioning workflow.

## Typical Workflow

1. Build or download the plugin `.jar` / `.json` artifact.
2. Reference the local file path in `delphix_database_plugin`.
3. Apply — DCT uploads the plugin to the target engine(s).
4. The plugin becomes available for linking AppData sources.

## Notes

- Plugin uploads are one-directional; there is no `Read` that can diff plugin versions — plan accordingly for upgrades.
- Available since provider v4.2.0.
