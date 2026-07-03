# Context:
<!--
A clear description of the high level effort that this pull request is
a part of. Anyone in the organization can see this change and may not
have the same context as you. Replace this comment with your text.
-->

# Problem:
<!--
A clear description of the problem. The problem statement should be
written in terms of a specific symptom that affects users or the
business. The problem statement should not be written in terms of the
solution. Replace this comment with your text.
-->

# Solution:
<!--
A clear description of the high-level solution you have chosen. If
there were other possible solutions that you considered and rejected,
please mention those as well. Please do not describe implementation
details when writing about the solution, those should go into the
implementation section. Replace this comment with your text.
-->

# Acceptance Criteria:
<!--
REQUIRED. State the criteria as verifiable, assertion-style statements
so a reviewer (or an automated agent) can confirm the change is done.
Prefer Given / When / Then, or "<subject> must <observable outcome>".
Replace this comment with your criteria. Example:

- Given a delphix_engine_configuration with a GCS object store, when
  terraform apply runs, then the engine is configured and the job
  reaches COMPLETED.
- Given an invalid bucket name, when apply runs, then the provider
  returns a validation error and no job is started.
- Overall CI coverage gate must fail the build when total coverage
  drops below COVERAGE_THRESHOLD.
-->

# Testing
<!--
Describe explicitly how this change was tested. Were we able to
recreate the issue? Were we able to write a test case that failed
before that passes now? Replace this comment with your text.
-->

## PR Checklist
<!-- All boxes must be checked before this PR can be merged. -->
- [ ] The **Problem**, **Solution**, **Acceptance Criteria**, and **Testing** sections above are filled in (placeholder comments replaced).
- [ ] Acceptance Criteria are written as verifiable assertions (Given/When/Then or "must" statements).
- [ ] New/changed code has unit tests; patch coverage meets the CI threshold (`PATCH_COVERAGE_THRESHOLD`).
- [ ] Commits are signed (GPG or SSH).

<!--
# Implementation:
The implementation details of the solution.
-->
<!--
# Notes To Reviewers:
Any extra information a reviewer may need to know before reviewing
your change. For example here you might want to describe which files
should be looked at first or which files are auto-generated.
-->
<!--
# Deployment Plan:
Some changes get more complicated and may need changes in multiple
repositories or may require infrastructure changes. Describe how these
changes will be smoothly deployed.
-->
<!--
# Future work:
A description of what follow up work is explicitly not being done in
this change.
-->
<!--
# Bonus:
A description of extra problems you've solved in this change. Did you
reformat an unrelated docstring? Point it out here
-->
