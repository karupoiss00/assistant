# Contract: Workspace Configuration

## Purpose

Define the repository-managed configuration expectations for locating and
governing the production workspace.

## Required Configuration Inputs

- `workspace_root_path`: explicit path to the production Obsidian-compatible
  workspace repository
- `log_retention_days`: minimum retention for assistant audit logs
- `default_sphere_priority_order`: default ordering used for future rituals and
  summaries
- `normalization_mode`: whether normalization is report-only or may apply safe
  structural additions
- `confirmation_required_for_obsidian_changes`: explicit guard for `.obsidian/`
  rewrites or replacements

## Contract Rules

- The workspace root must be explicit and configurable; it may not be inferred
  from `.workspace/`.
- Configuration must live in repository-managed files, not hidden runtime state.
- Configuration changes that only adjust future assistant behavior are
  repository edits, not workspace normalization events.
- Any prompt or automation that performs normalization must consume the same
  action classifications defined in
  [normalization-policy.md](./normalization-policy.md).

## Expected Repository Touchpoints

- `config/` stores assistant behavior and workspace-location configuration.
- `schemas/` stores formal definitions for any future config document and
  normalization report format.
- `prompts/` stores the prompt contracts that must not bypass confirmation
  boundaries.

## Validation Expectations

- Operators can inspect the configured workspace root without reading code.
- Operators can verify that log retention is at least 30 days.
- Operators can determine whether a given normalization flow is allowed to apply
  structural changes automatically or must stay report-only.
- Operators can verify that required artifacts remain manually readable and
  editable through Obsidian.
