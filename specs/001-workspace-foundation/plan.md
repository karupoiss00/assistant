# Implementation Plan: Workspace Foundation

**Branch**: `001-workspace-foundation` | **Date**: 2026-06-08 | **Spec**: [spec.md](./spec.md)

**Input**: Feature specification from `/specs/001-workspace-foundation/spec.md`

## Summary

Define the `v1` contract for an Obsidian-compatible workspace repository by
formalizing the reference example in `.workspace`, documenting configurable
workspace-root resolution, and specifying a safe normalization policy that can
add missing structural artifacts without rewriting user meaning. The technical
approach is documentation-first: produce explicit repository contracts,
normalization rules, config/logging expectations, and validation scenarios
before any automation writes to a production workspace.

## Technical Context

**Language/Version**: Markdown documentation and repository contracts; shell
automation expected to remain Bash 3.2+ compatible for local workflow scripts

**Primary Dependencies**: Spec Kit workflow in `.specify/`; Obsidian-compatible
Markdown repository structure; Git history for audit and reversibility

**Storage**: File-based storage inside the production workspace repository and
this repository's versioned docs/config directories

**Testing**: Specification checklist review, contract walkthroughs, and manual
validation scenarios captured in `quickstart.md`; future state-changing
automation will require integration tests before implementation is complete

**Target Platform**: Local macOS/Linux operator environment managing an
Obsidian-compatible Git repository

**Project Type**: Repository contract and automation-planning feature for a
stateful assistant workspace

**Performance Goals**: Normalization planning should stay inspectable by a
single operator in under 15 minutes and produce a deterministic path-by-path
change report for every run

**Constraints**: Production workspace path must be configurable; no durable
state may live in `runtime/`; normalization must preserve existing user notes,
links, attachments, and `.obsidian` settings unless the user confirms a
meaning-changing mutation; logs must remain reviewable for at least 30 days

**Scale/Scope**: One production workspace repository, five required `v1`
spheres, one shared template sphere, repository-managed assistant config/log
contracts, and enough structure to support later capture/review/linking stages

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [x] Authoritative repository files and directories are identified; no durable
      state is delegated to `runtime/` or undocumented storage.
- [x] Each planned mutation is classified as `auto-allowed`, `proposal-only`,
      or `requires-confirmation`, with explicit rationale for any automatic
      write.
- [x] The plan states which ritual(s) or planning horizon(s) the feature
      improves and how it preserves cross-sphere balance.
- [x] Logging, retention, and operator audit paths are defined for every
      state-changing flow and failure path.
- [x] Required contract updates are listed for schemas, prompts, config, and
      repository note formats.
- [x] Validation covers both user-visible behavior and resulting repository
      mutations.

## Project Structure

### Documentation (this feature)

```text
specs/001-workspace-foundation/
├── checklists/
│   ├── requirements.md
│   └── implementation.md
├── plan.md
├── research.md
├── data-model.md
├── quickstart.md
├── contracts/
│   ├── normalization-policy.md
│   ├── workspace-config-contract.md
│   └── workspace-filesystem-contract.md
└── tasks.md
```

### Source Code (repository root)

```text
.workspace/                  # Reference workspace example only
config/
├── README.md
├── defaults.md
└── workspace.md
logs/
├── README.md
├── .gitkeep
└── action-log.md
prompts/
├── README.md
├── .gitkeep
└── normalization.md
schemas/
├── README.md
└── workspace-config.schema.md
specs/
└── 001-workspace-foundation/
runtime/                     # Ephemeral only; never a source of truth
AGENTS.md
```

**Structure Decision**: This feature is documentation- and contract-centric.
Implementation planning therefore anchors on repository paths that already exist
for durable contracts (`specs/`, `config/`, `logs/`, `prompts/`, `schemas/`)
plus the `.workspace/` reference example. No application `src/` tree is needed
for this stage because the outcome is the written contract that later automation
must implement.

### Target Artifact Inventory

- `specs/001-workspace-foundation/checklists/implementation.md`: execution-time
  validation checklist for implemented contract artifacts
- `config/workspace.md`: operator-facing summary of source-of-truth, planning
  chain, and Obsidian editability rules
- `logs/action-log.md`: repository-backed audit log contract
- `schemas/workspace-config.schema.md`: schema documentation for workspace
  config and normalization reports
- `prompts/normalization.md`: prompt contract for future normalization flows

## Phase 0: Research

### Research Goals

- Confirm the repository contract should treat `.workspace/` as a reference
  example rather than the fixed production path.
- Define the safest normalization policy that preserves user-authored notes,
  links, attachments, and `.obsidian` state.
- Decide how repository config, prompts, schemas, and logs should participate
  in the workspace contract without becoming hidden runtime state.

### Research Output

- [research.md](./research.md) resolves the technical choices for workspace-root
  configuration, normalization semantics, logging retention, and repository
  contract boundaries.

## Phase 1: Design & Contracts

### Planned Artifacts

- [data-model.md](./data-model.md): entities for workspace root, sphere,
  artifact templates, normalization actions, logs, and config.
- [checklists/implementation.md](./checklists/implementation.md): execution-time
  validation checklist for contract completeness, normalization safety, custom
  sphere preservation, auditability, and Obsidian readability.
- [contracts/workspace-filesystem-contract.md](./contracts/workspace-filesystem-contract.md):
  authoritative repository layout and required artifacts.
- [contracts/workspace-config-contract.md](./contracts/workspace-config-contract.md):
  contract for resolving the production workspace path and repository-managed
  assistant settings.
- [contracts/normalization-policy.md](./contracts/normalization-policy.md):
  write-boundary rules, action classification, and audit output format.
- [quickstart.md](./quickstart.md): validation walkthroughs for contract review,
  safe normalization, and audit behavior.
- [config/workspace.md](/Users/rostislav.glizerin/projects/assistant/config/workspace.md):
  repository-managed summary of workspace-root, source-of-truth, and
  Obsidian-facing contract rules.
- [logs/action-log.md](/Users/rostislav.glizerin/projects/assistant/logs/action-log.md):
  repository-backed audit log contract and report expectations.
- [schemas/workspace-config.schema.md](/Users/rostislav.glizerin/projects/assistant/schemas/workspace-config.schema.md):
  schema contract for workspace-root resolution and normalization reports.
- [prompts/normalization.md](/Users/rostislav.glizerin/projects/assistant/prompts/normalization.md):
  prompt-side confirmation and write-boundary contract for future automation.

### Planned Mutations

- `auto-allowed`: Create missing required directories and missing required note
  stubs; create dedicated config/log locations when absent; append missing
  structural sections that do not overwrite user-authored content.
- `proposal-only`: Recommend renames for non-standard sphere naming, suggest
  note consolidation, and propose new semantic links across artifacts.
- `requires-confirmation`: Delete, archive, rename, merge, move, or overwrite
  existing user notes or `.obsidian` settings; any change that alters current
  meaning or operator working layout.

### Ritual Coverage

- Supports monthly and longer-horizon planning by fixing the repository
  structure for the chain `Операционная задача -> Цель месяца -> Годовая цель ->
  Цель на 5 лет`.
- Preserves cross-sphere balance by making all five `v1` spheres first-class and
  forbidding hidden prioritization through undocumented automation.

### Logging and Audit

- Authoritative audit artifacts live in the production workspace repository and
  must remain readable from Obsidian plus Git history.
- Minimum retention is 30 days for detected contract violations, created
  structural elements, proposals not auto-applied, confirmed meaning-changing
  mutations, and normalization errors.
- Every normalization run must emit a path-by-path report that distinguishes
  automatic changes from proposals and confirmed changes.

### Required Contract Updates Before Implementation

- `config/`: add a repository-managed contract for production workspace root,
  logging retention, and future assistant behavior settings.
- `logs/`: define a repository-readable action-log format and retention rules.
- `schemas/`: define schemas for normalization reports and any future config
  document used to locate the workspace.
- `prompts/`: any later prompt that triggers normalization must reference the
  same action classifications and confirmation boundaries.
- Workspace note formats: preserve and document required sphere notes, shared
  notes, the planning chain `Операционная задача -> Цель месяца -> Годовая цель
  -> Цель на 5 лет`, manual readability/editability through Obsidian, and
  `.obsidian` expectations without encoding hidden assumptions in runtime logic.

### Validation Strategy

- Review the contract documents against the reference `.workspace/` tree to
  confirm all required `v1` artifacts, planning-chain anchor points, and
  Obsidian readability/editability rules are covered.
- Walk through normalization scenarios that add missing structure while leaving
  existing user content untouched, including preservation of additional
  user-defined top-level spheres.
- Verify each scenario produces a user-readable change report and references
  repository-backed logs.
- Re-check constitution compliance after design artifacts are written.

## Post-Design Constitution Check

- [x] Authoritative repository files and directories are identified in the plan,
      data model, and filesystem contract.
- [x] Planned mutations are classified in the plan and normalization policy.
- [x] Ritual support and cross-sphere balance are explicit in the plan and
      quickstart scenarios.
- [x] Logging and retention are defined in the plan, research decisions, and
      normalization policy.
- [x] Contract updates for config, prompts, schemas, logs, and note formats are
      listed before implementation.
- [x] Validation includes user-visible outcomes and repository mutation audit
      checks.

## Complexity Tracking

No constitution violations require justification for this stage.
