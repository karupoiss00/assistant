# Tasks: Workspace Foundation

**Input**: Design documents from `/specs/001-workspace-foundation/`

**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: This feature defines repository-backed state, normalization rules,
and audit requirements. Validation artifacts are therefore mandatory and are
included as checklist-based contract tests and quickstart walkthrough updates.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g. `US1`, `US2`, `US3`)
- Include exact file paths in descriptions

## Path Conventions

- Documentation and contracts for this feature live in `specs/001-workspace-foundation/`
- Repository-level durable contracts live in `config/`, `logs/`, `prompts/`, and `schemas/`
- Reference examples stay in `.workspace/` and must not be treated as the production root

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Establish repo-level files that later story tasks will update

- [X] T001 Create a target artifact inventory section in specs/001-workspace-foundation/plan.md
- [X] T002 [P] Create a workspace contract validation checklist in specs/001-workspace-foundation/checklists/implementation.md
- [X] T003 [P] Create a repo-level contract stub for workspace configuration in config/workspace.md
- [X] T004 [P] Create a repo-level audit log contract stub in logs/action-log.md
- [X] T005 [P] Create schema stubs for workspace config and normalization reports in schemas/workspace-config.schema.md
- [X] T006 [P] Create a prompt contract stub for future normalization flows in prompts/normalization.md

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core contract infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T007 Consolidate repository-first source-of-truth rules in config/workspace.md
- [X] T008 [P] Define normalization report fields and retention rules in logs/action-log.md
- [X] T009 [P] Define workspace root and normalization mode schema fields in schemas/workspace-config.schema.md
- [X] T010 [P] Define confirmation-boundary requirements for prompts in prompts/normalization.md
- [X] T011 Define autonomy classes and write-boundary glossary in specs/001-workspace-foundation/contracts/normalization-policy.md
- [X] T012 Define repository-backed validation steps that map to quickstart scenarios in specs/001-workspace-foundation/checklists/implementation.md

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Understand Workspace Contract (Priority: P1) 🎯 MVP

**Goal**: Publish a complete, inspectable workspace contract that explains required spheres, shared artifacts, template semantics, and `.obsidian` role

**Independent Test**: A reviewer can read the filesystem/config contract docs and identify required `v1` spheres, shared artifacts, configurable workspace root, and `.obsidian` preservation rules without opening source code

### Tests for User Story 1 ⚠️

- [X] T013 [P] [US1] Add a contract review checklist for workspace structure, planning-chain anchors, and Obsidian editability in specs/001-workspace-foundation/checklists/implementation.md
- [X] T014 [P] [US1] Add a quickstart validation scenario for configurable workspace-root resolution in specs/001-workspace-foundation/quickstart.md
- [X] T015 [P] [US1] Add a contract consistency review section covering `.workspace` reference usage in specs/001-workspace-foundation/research.md

### Implementation for User Story 1

- [X] T016 [P] [US1] Finalize required sphere, shared-artifact, planning-chain, and Obsidian readability rules in specs/001-workspace-foundation/contracts/workspace-filesystem-contract.md
- [X] T017 [P] [US1] Finalize workspace-root configuration rules in specs/001-workspace-foundation/contracts/workspace-config-contract.md
- [X] T018 [P] [US1] Document workspace-root, shared-artifact, and planning-chain entities in specs/001-workspace-foundation/data-model.md
- [X] T019 [US1] Publish the repository-managed workspace contract summary, including manual Obsidian editability rules, in config/workspace.md
- [X] T020 [US1] Cross-link plan, contracts, and quickstart for user-readable contract navigation in specs/001-workspace-foundation/plan.md

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - Normalize Existing Workspace Safely (Priority: P2)

**Goal**: Define a safe normalization policy that adds missing required structure without overwriting user-authored content

**Independent Test**: A reviewer can walk through a non-compliant workspace scenario and determine which changes are additive, which preserve existing notes, and which are blocked from automatic execution

### Tests for User Story 2 ⚠️

- [X] T021 [P] [US2] Add a normalization walkthrough checklist for missing required artifacts and preservation of custom top-level spheres in specs/001-workspace-foundation/checklists/implementation.md
- [X] T022 [P] [US2] Add a quickstart scenario for preserving user notes, `.obsidian` files, and custom spheres during normalization in specs/001-workspace-foundation/quickstart.md
- [X] T023 [P] [US2] Add a report-format review checklist for normalization outcomes in logs/action-log.md

### Implementation for User Story 2

- [X] T024 [P] [US2] Finalize additive-only normalization rules, including preservation of custom spheres, in specs/001-workspace-foundation/contracts/normalization-policy.md
- [X] T025 [P] [US2] Document sphere normalization states and action transitions in specs/001-workspace-foundation/data-model.md
- [X] T026 [P] [US2] Define normalization mode and confirmation flags in schemas/workspace-config.schema.md
- [X] T027 [US2] Document safe-structure creation rules and forbidden rewrites in prompts/normalization.md
- [X] T028 [US2] Publish the operator-facing normalization policy summary in config/workspace.md

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: User Story 3 - Define Source of Truth and Write Boundaries (Priority: P3)

**Goal**: Make source-of-truth boundaries, action classes, and audit expectations explicit for users and future automation

**Independent Test**: A reviewer can classify any planned workspace mutation as `auto-allowed`, `proposal-only`, or `requires-confirmation` and can identify where the resulting audit trail must live

### Tests for User Story 3 ⚠️

- [X] T029 [P] [US3] Add an autonomy-classification review checklist in specs/001-workspace-foundation/checklists/implementation.md
- [X] T030 [P] [US3] Add a quickstart scenario for audit-log inspection and 30-day retention in specs/001-workspace-foundation/quickstart.md
- [X] T031 [P] [US3] Add a prompt-governance review section for confirmation-gated actions in prompts/normalization.md

### Implementation for User Story 3

- [X] T032 [P] [US3] Finalize action classes, preservation rules, and report contract in specs/001-workspace-foundation/contracts/normalization-policy.md
- [X] T033 [P] [US3] Finalize audit-log entity fields and retention requirements in specs/001-workspace-foundation/data-model.md
- [X] T034 [P] [US3] Publish repository-backed audit record expectations in logs/action-log.md
- [X] T035 [P] [US3] Define normalization report schema fields and examples in schemas/workspace-config.schema.md
- [X] T036 [US3] Publish source-of-truth and confirmation-boundary guidance in config/workspace.md

**Checkpoint**: All user stories should now be independently functional

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [X] T037 [P] Align feature spec status and cross-references with implemented contract artifacts in specs/001-workspace-foundation/spec.md
- [X] T038 [P] Review retention, audit visibility, and repository-first wording across config/workspace.md, logs/action-log.md, and prompts/normalization.md
- [X] T039 [P] Run the end-to-end contract walkthrough and record results in specs/001-workspace-foundation/checklists/implementation.md
- [X] T040 Validate all quickstart scenarios, including planning-chain coverage and Obsidian editability, against the final artifact set in specs/001-workspace-foundation/quickstart.md

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3+)**: All depend on Foundational phase completion
  - User stories can then proceed in parallel
  - Or sequentially in priority order (`P1` → `P2` → `P3`)
- **Polish (Phase 6)**: Depends on all desired user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2) - No dependencies on other stories
- **User Story 2 (P2)**: Can start after Foundational (Phase 2) - Builds on shared contract files but remains independently testable
- **User Story 3 (P3)**: Can start after Foundational (Phase 2) - Builds on shared contract files but remains independently testable

### Within Each User Story

- Validation checklist and quickstart tasks come before implementation updates
- Shared entity and contract definitions come before repo-level summaries
- Repo-level summaries come before final cross-linking and validation

### Parallel Opportunities

- Setup tasks marked `[P]` can run in parallel
- Foundational tasks marked `[P]` can run in parallel
- Once Foundational completes, the three user stories can proceed in parallel if different people coordinate file ownership
- Test and checklist tasks marked `[P]` within a story can run in parallel
- Contract/data-model updates marked `[P]` within a story can run in parallel before repo-level summary tasks

---

## Parallel Example: User Story 1

```bash
Task: "Add a contract review checklist for workspace structure, planning-chain anchors, and Obsidian editability in specs/001-workspace-foundation/checklists/implementation.md"
Task: "Add a quickstart validation scenario for configurable workspace-root resolution in specs/001-workspace-foundation/quickstart.md"
Task: "Add a contract consistency review section covering .workspace reference usage in specs/001-workspace-foundation/research.md"
```

```bash
Task: "Finalize required sphere, shared-artifact, planning-chain, and Obsidian readability rules in specs/001-workspace-foundation/contracts/workspace-filesystem-contract.md"
Task: "Finalize workspace-root configuration rules in specs/001-workspace-foundation/contracts/workspace-config-contract.md"
Task: "Document workspace-root, shared-artifact, and planning-chain entities in specs/001-workspace-foundation/data-model.md"
```

---

## Parallel Example: User Story 2

```bash
Task: "Add a normalization walkthrough checklist for missing required artifacts and preservation of custom top-level spheres in specs/001-workspace-foundation/checklists/implementation.md"
Task: "Add a quickstart scenario for preserving user notes, .obsidian files, and custom spheres during normalization in specs/001-workspace-foundation/quickstart.md"
Task: "Add a report-format review checklist for normalization outcomes in logs/action-log.md"
```

```bash
Task: "Finalize additive-only normalization rules, including preservation of custom spheres, in specs/001-workspace-foundation/contracts/normalization-policy.md"
Task: "Document sphere normalization states and action transitions in specs/001-workspace-foundation/data-model.md"
Task: "Define normalization mode and confirmation flags in schemas/workspace-config.schema.md"
```

---

## Parallel Example: User Story 3

```bash
Task: "Add an autonomy-classification review checklist in specs/001-workspace-foundation/checklists/implementation.md"
Task: "Add a quickstart scenario for audit-log inspection and 30-day retention in specs/001-workspace-foundation/quickstart.md"
Task: "Add a prompt-governance review section for confirmation-gated actions in prompts/normalization.md"
```

```bash
Task: "Finalize action classes, preservation rules, and report contract in specs/001-workspace-foundation/contracts/normalization-policy.md"
Task: "Finalize audit-log entity fields and retention requirements in specs/001-workspace-foundation/data-model.md"
Task: "Publish repository-backed audit record expectations in logs/action-log.md"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational
3. Complete Phase 3: User Story 1
4. Validate the workspace contract with the checklist and quickstart artifacts
5. Stop and review before moving to normalization semantics

### Incremental Delivery

1. Complete Setup + Foundational to establish repo-level contract files
2. Deliver User Story 1 and validate contract readability
3. Deliver User Story 2 and validate safe normalization semantics
4. Deliver User Story 3 and validate auditability plus confirmation boundaries
5. Finish with cross-cutting consistency review

### Parallel Team Strategy

1. One person completes Setup + Foundational
2. After that:
   - Developer A: filesystem/config contract updates for US1
   - Developer B: normalization policy and schema updates for US2
   - Developer C: audit/log/prompt updates for US3
3. Rejoin for Polish phase and final walkthrough
