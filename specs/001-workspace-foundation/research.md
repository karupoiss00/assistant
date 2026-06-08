# Research: Workspace Foundation

## Decision 1: Treat `.workspace/` as a reference example, not the production root

- **Decision**: Use `.workspace/` only as the source example for `v1` contract
  discovery. The production workspace root must be resolved from repository
  configuration rather than hard-coded to `.workspace/`.
- **Rationale**: The feature specification explicitly requires portability to an
  arbitrary Obsidian Git repository. Hard-coding `.workspace/` would violate the
  contract before implementation begins.
- **Alternatives considered**:
  - Bind `v1` to `.workspace/` permanently: rejected because it couples the
    product to this repo's local reference layout.
  - Discover the workspace root heuristically at runtime: rejected because it
    introduces hidden behavior and ambiguity about the source of truth.

## Decision 2: Normalize by additive structure only unless meaning changes

- **Decision**: Define normalization as additive by default: create missing
  required directories, note stubs, and structural sections, but never delete,
  rename, archive, move, or overwrite existing user-authored content without
  confirmation.
- **Rationale**: The reference workspace already contains meaningful notes and
  `.obsidian` state. The safest contract is one where structural compliance can
  improve without silent loss of context.
- **Alternatives considered**:
  - Rebuild non-conforming spheres from the template: rejected because it risks
    overwriting user meaning.
  - Auto-rename custom notes into contract names: rejected because name changes
    alter user navigation and implied semantics.

## Decision 3: Preserve `.obsidian/` as user-owned working state

- **Decision**: Treat `.obsidian/` as part of the authoritative workspace that
  must be preserved during normalization. Replacement or rewrite of existing
  `.obsidian` settings requires confirmation.
- **Rationale**: The reference example shows real client settings and layout
  files such as `workspace.json`. These affect user workflow directly and are
  therefore meaning-changing.
- **Alternatives considered**:
  - Ignore `.obsidian/` entirely: rejected because the spec requires its role to
    be documented and preserved.
  - Fully standardize `.obsidian/` from a template: rejected because it would
    silently change the operator's working environment.

## Decision 4: Keep assistant config, prompts, schemas, and logs explicit in-repo

- **Decision**: The workspace contract must explicitly point to repository
  locations for config, prompts, schemas, and logs. Future automation must read
  and write through those documented contracts instead of hidden runtime files.
- **Rationale**: The constitution requires repository-first truth and explicit
  contracts. Existing repo directories `config/`, `logs/`, `prompts/`, and
  `schemas/` already establish the intended separation of responsibilities.
- **Alternatives considered**:
  - Store config and logs only in runtime state: rejected because durable state
    would become opaque and unauditable.
  - Embed all behavior assumptions in code only: rejected because contracts
    would drift and become hard to inspect.

## Decision 5: Audit every normalization run with a path-by-path report

- **Decision**: Require each normalization flow to emit a user-readable report
  that lists created or amended structural elements, proposals not executed,
  confirmation-gated actions, and any errors, with at least 30 days of
  repository-backed retention.
- **Rationale**: Trust in autonomous structural changes depends on reversibility
  and traceability. Git history alone is not enough because the user also needs
  a semantic summary of what changed and why.
- **Alternatives considered**:
  - Rely only on Git diffs: rejected because diffs do not classify intent or
    separate auto-allowed from proposed actions.
  - Keep transient execution logs outside the repository: rejected because they
    would not satisfy the source-of-truth and audit requirements.

## Contract Consistency Review

- `.workspace/` remains a design-time reference example and does not acquire
  production write authority anywhere in the contract set.
- The planning chain `Операционная задача -> Цель месяца -> Годовая цель ->
  Цель на 5 лет` is part of the documented workspace contract even though later
  stages will refine link syntax and orchestration.
- Required artifacts remain manually readable and editable through Obsidian, so
  Telegram does not become the only durable interface.
- Additional user-defined top-level spheres remain valid user content and are
  preserved during normalization unless the user approves a meaning-changing
  reorganization.
