# docs/guides/

Hand-authored guides published to the Terraform Registry under the **Guides** section.

## Files

| File | Purpose |
|---|---|
| [provider_guide.md](provider_guide.md) | Getting-started overview — points users to the `examples/` directory and GitHub Issues for requests |

## Adding a Guide

1. Create a new `.md` file here with a descriptive name (e.g., `oracle_dsource_guide.md`).
2. Add a frontmatter block at the top that the Registry requires:
   ```markdown
   ---
   page_title: "Guide Title"
   subcategory: ""
   description: "Short description for Registry search"
   ---
   ```
3. Link to the relevant example under [../../examples/](../../examples/).

Unlike `resources/`, these files are **not** auto-generated and must be maintained manually.
