# Contract: Normalization Policy

## Purpose

Define how the assistant may bring an existing workspace into `v1` contract
compliance without silently changing user meaning.

## Action Classes

### Auto-Allowed

These actions may execute without prior confirmation:

- create a missing required sphere directory
- create a missing required note stub
- create a missing `Артефакты/` directory
- create repository-backed locations for assistant config or logs when they are
  part of the documented contract
- append missing structural sections when the append does not rewrite or remove
  existing user-authored content

### Proposal-Only

These actions must be reported as suggestions unless the user explicitly chooses
to continue:

- propose renaming non-standard sphere or note names to the contract names
- propose consolidating overlapping notes
- propose adding semantic links between notes beyond required structural anchors
- propose standardizing content that differs from the reference template but is
  still user-authored

### Requires-Confirmation

These actions may not execute without explicit confirmation:

- delete, archive, rename, move, or merge existing notes or directories
- overwrite existing note content
- rewrite or replace `.obsidian/` settings or layout files
- relocate user files between spheres
- change any content whose effect is to alter current user meaning or working
  environment

## Preservation Rules

- Existing user notes remain authoritative even when they diverge from the
  reference template.
- Existing links, attachments, and artifact references must survive
  normalization unchanged unless the user confirms a meaning-changing action.
- `.obsidian/` is preserved as user-owned client state when present.
- Additional user-defined top-level spheres are preserved even when they are not
  part of the `v1` minimum structure.
- Required artifacts must remain manually readable and editable through
  Obsidian after structural normalization.

## Normalization Report Contract

Each normalization run must produce a repository-backed report containing:

- resolved production workspace root
- timestamp and run identifier
- detected contract violations
- executed `auto-allowed` actions with target paths
- proposed but unexecuted actions with rationale
- confirmation-gated actions and their final status
- errors or skipped paths

## Audit and Retention

- Every report must remain accessible from the repository for at least 30 days.
- Every executed or proposed action must map to a stable audit entry.
- Users must be able to distinguish automatic structural additions from proposed
  or confirmed meaning-changing actions by reading the report alone.
