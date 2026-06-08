# Tasks: Telegram Capture

**Input**: Design documents from `/specs/002-telegram-capture/`

**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: This feature introduces repository-backed capture state, Telegram-facing statuses, webhook runtime behavior, retention requirements, and temporary runtime cleanup rules. Validation artifacts are therefore mandatory and are included as checklist-based contract tests, webhook smoke checks, and quickstart walkthrough updates.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g. `US1`, `US2`, `US3`)
- Include exact file paths in descriptions

## Path Conventions

- Feature-specific documentation and contracts live in `specs/002-telegram-capture/`
- Repository-level durable contracts live in `config/`, `logs/`, `prompts/`, `runtime/`, and `schemas/`
- Runtime implementation tasks in this feature target:
  - `go.mod`
  - `cmd/telegram-capture/main.go`
  - `internal/config/`
  - `internal/telegram/`
  - `internal/capture/`
  - `internal/workspace/`
  - `internal/audit/`
  - `internal/transcribe/`
- Temporary voice artifacts may exist only under `runtime/` and must never become authoritative repository state

## Phase 1: Contract Cleanup (Shared Infrastructure)

**Purpose**: Establish repo-level artifacts that later story tasks will refine

- [x] T001 Create a Telegram capture validation checklist in specs/002-telegram-capture/checklists/implementation.md
- [x] T002 [P] Create a Stage 2 capture contract section stub in config/workspace.md
- [x] T003 [P] Create a Telegram capture audit contract stub in logs/action-log.md
- [x] T004 [P] Create capture schema stubs for daily inbox entries and audit events in schemas/workspace-config.schema.md
- [x] T005 [P] Create a prompt contract stub for Telegram capture confirmations and downstream classification handoff in prompts/normalization.md
- [x] T006 [P] Create a temporary voice artifact lifecycle section in runtime/README.md

---

## Phase 2: Foundational Contracts (Blocking Prerequisites)

**Purpose**: Core capture infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T007 Consolidate repository-first Telegram capture source-of-truth rules in config/workspace.md
- [x] T008 [P] Define capture event taxonomy, Telegram correlation fields, and 30-day retention rules in logs/action-log.md
- [x] T009 [P] Define daily inbox entry, capture audit, and retry-report schema fields in schemas/workspace-config.schema.md
- [x] T010 [P] Define auto-allowed, proposal-only, and requires-confirmation boundaries for capture follow-up actions in prompts/normalization.md
- [x] T011 [P] Define runtime-only download, cleanup, and deletion guarantees for temporary voice artifacts in runtime/README.md
- [x] T012 Define repository mutation checks, webhook smoke validation steps, and quickstart-backed validation steps in specs/002-telegram-capture/checklists/implementation.md
- [x] T013 Define the minimum Go webhook runtime contract, including Telegram delivery, user reply, repository write path, and service boundaries, in specs/002-telegram-capture/plan.md

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: Runtime Bootstrap (Shared Code Foundation)

**Purpose**: Introduce the concrete Go runtime skeleton required by all user stories

- [x] T014 Create `go.mod` for the Stage 2 runtime
- [x] T015 [P] Create `cmd/telegram-capture/main.go` with HTTP server bootstrap
- [x] T016 [P] Create `internal/config/` for assistant config loading and validation
- [x] T017 [P] Create `internal/workspace/` for daily inbox writes and path resolution
- [x] T018 [P] Create `internal/audit/` for append-only capture audit writes
- [x] T019 [P] Create `internal/telegram/` for webhook payload parsing and Telegram replies
- [x] T020 [P] Create `internal/transcribe/` for speech-to-text client abstraction
- [x] T021 Create `internal/capture/` to orchestrate text and voice capture flows

**Checkpoint**: Runtime bootstrap exists and user story work can target concrete code paths

---

## Phase 4: User Story 1 - Capture Text Reliably (Priority: P1) 🎯 MVP

**Goal**: Publish a complete text-capture contract so Telegram text inputs reliably land in a daily inbox file plus audit logs without requiring classification

**Independent Test**: A reviewer can trace one text Telegram update from accepted input to a daily inbox entry and a separate audit record using only the repository contracts and validation artifacts

### Tests for User Story 1 ⚠️

- [x] T022 [P] [US1] Add a checklist section for text capture durability and one-update-to-one-entry mapping in specs/002-telegram-capture/checklists/implementation.md
- [x] T023 [P] [US1] Add a quickstart scenario for text capture confirmation and repository verification in specs/002-telegram-capture/quickstart.md
- [x] T024 [P] [US1] Add a webhook smoke-validation scenario for text capture intake in specs/002-telegram-capture/quickstart.md
- [x] T025 [P] [US1] Add a lightweight latency-validation scenario for text capture confirmation under the 10-second target in specs/002-telegram-capture/quickstart.md

### Implementation for User Story 1

- [x] T026 [P] [US1] Finalize daily inbox file rules, entry identity, and Obsidian readability expectations in specs/002-telegram-capture/contracts/telegram-capture-filesystem-contract.md
- [x] T027 [P] [US1] Finalize accepted text input metadata, webhook intake semantics, and saved-status contract in specs/002-telegram-capture/contracts/telegram-capture-telegram-contract.md
- [x] T028 [P] [US1] Document text message, daily inbox file, and capture buffer entry entities in specs/002-telegram-capture/data-model.md
- [x] T029 [P] [US1] Define text capture examples and mandatory text-entry fields in schemas/workspace-config.schema.md
- [ ] T030 [P] [US1] Implement text webhook handler in `internal/telegram/` and `internal/capture/`
- [ ] T031 [P] [US1] Implement daily inbox append and audit append in `internal/workspace/` and `internal/audit/`
- [x] T032 [US1] Publish the operator-facing text capture contract summary in config/workspace.md
- [ ] T033 [US1] Wire text capture bootstrap through `cmd/telegram-capture/main.go`

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 5: User Story 2 - Capture Voice with Transcript (Priority: P2)

**Goal**: Define a reliable voice-capture flow where transcription and basic understanding happen before user-visible success, while temporary audio is cleaned up afterward

**Independent Test**: A reviewer can walk through a voice message flow and confirm that transcript, metadata, and audit records become durable while the temporary voice file is deleted after processing or failure handling

### Tests for User Story 2 ⚠️

- [x] T034 [P] [US2] Add a checklist section for voice transcription confirmation, failure handling, and temporary-file cleanup in specs/002-telegram-capture/checklists/implementation.md
- [x] T035 [P] [US2] Add a quickstart scenario for voice capture success, failure-aware status, and cleanup verification in specs/002-telegram-capture/quickstart.md
- [x] T036 [P] [US2] Add a voice-processing audit review section covering transcription attempts and cleanup events in logs/action-log.md
- [x] T037 [P] [US2] Add a webhook smoke-validation scenario for voice capture intake in specs/002-telegram-capture/quickstart.md
- [x] T038 [P] [US2] Add a lightweight latency-validation scenario for voice confirmation under the 10-second target in specs/002-telegram-capture/quickstart.md

### Implementation for User Story 2

- [x] T039 [P] [US2] Finalize voice processing, user-visible confirmation, and cleanup rules in specs/002-telegram-capture/contracts/telegram-capture-processing-contract.md
- [x] T040 [P] [US2] Finalize voice input metadata and failure-aware user statuses in specs/002-telegram-capture/contracts/telegram-capture-telegram-contract.md
- [x] T041 [P] [US2] Document voice transcript, transcription attempt, and temporary voice artifact entities in specs/002-telegram-capture/data-model.md
- [x] T042 [P] [US2] Define voice transcript, processing-status, and cleanup-report schema fields in schemas/workspace-config.schema.md
- [x] T043 [P] [US2] Publish repository-backed transcription, failure, and cleanup event expectations in logs/action-log.md
- [ ] T044 [P] [US2] Implement voice transcription flow in `internal/capture/`, `internal/transcribe/`, and `internal/workspace/`
- [ ] T045 [P] [US2] Implement temporary artifact cleanup in `internal/capture/` and `runtime/` handling
- [x] T046 [US2] Publish user-facing confirmation and failure wording guidance for voice capture in prompts/normalization.md
- [x] T047 [US2] Publish the operator-facing voice capture and temporary-storage policy summary in config/workspace.md

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 6: User Story 3 - Preserve Queue Integrity and Auditability (Priority: P3)

**Goal**: Make rapid-message handling, repeats, retries, and auditability explicit so Stage 2 remains trustworthy under real Telegram usage

**Independent Test**: A reviewer can inspect the contracts and determine how five rapid messages, a repeated Telegram update, and a retry after failure are preserved as separate observable capture facts with traceable logs

### Tests for User Story 3 ⚠️

- [x] T048 [P] [US3] Add a checklist section for burst traffic, repeated updates, retry-safe processing, repository write failure handling, and missing Stage 1 dependency behavior in specs/002-telegram-capture/checklists/implementation.md
- [x] T049 [P] [US3] Add a quickstart scenario for rapid-message ordering, repeat visibility, retry validation, repository write failure behavior, and missing Stage 1 dependency behavior in specs/002-telegram-capture/quickstart.md
- [x] T050 [P] [US3] Add an audit retention and repeat-activity review section in logs/action-log.md

### Implementation for User Story 3

- [x] T051 [P] [US3] Finalize burst-traffic, repeat-update, and non-dedup posture in specs/002-telegram-capture/contracts/telegram-capture-telegram-contract.md
- [x] T052 [P] [US3] Finalize retry semantics, failure preservation, webhook failure responses, and reprocessing boundaries in specs/002-telegram-capture/contracts/telegram-capture-processing-contract.md
- [x] T053 [P] [US3] Document capture audit record, retry attempt, repeat-activity, and repository-write-failure entities in specs/002-telegram-capture/data-model.md
- [x] T054 [P] [US3] Define repeat-event, retry-request, burst-order, and write-failure schema examples in schemas/workspace-config.schema.md
- [x] T055 [P] [US3] Publish retry-safe audit record expectations, repeat correlation rules, and write-failure events in logs/action-log.md
- [ ] T056 [P] [US3] Implement burst, repeat, retry, and workspace failure handling in `internal/capture/` and `internal/audit/`
- [x] T057 [US3] Publish no-semantic-dedup and reprocessing guidance in config/workspace.md

**Checkpoint**: All user stories should now be independently functional

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [x] T058 [P] Align the clarified voice-storage, webhook runtime, and confirmation rules with the final wording in specs/002-telegram-capture/spec.md
- [ ] T059 [P] Review repository-first, retention, webhook, and cleanup wording across config/workspace.md, logs/action-log.md, prompts/normalization.md, and runtime/README.md
- [ ] T060 [P] Run the end-to-end contract and webhook walkthrough and record results in specs/002-telegram-capture/checklists/implementation.md
- [ ] T061 Validate all quickstart scenarios, including text durability, voice cleanup, burst handling, retry traceability, webhook smoke behavior, and latency checks, against the final artifact set in specs/002-telegram-capture/quickstart.md
- [ ] T062 [P] Record webhook-to-user-reply latency measurements for text and voice flows and evaluate the 95% / 10-second target in specs/002-telegram-capture/checklists/implementation.md

---

## Dependencies & Execution Order

### Phase Dependencies

- **Contract Cleanup (Phase 1)**: No dependencies - can start immediately
- **Foundational Contracts (Phase 2)**: Depends on Phase 1 completion - BLOCKS all user stories
- **Runtime Bootstrap (Phase 3)**: Depends on Foundational phase completion and BLOCKS code implementation for all user stories
- **User Stories (Phase 4+)**: All depend on Runtime Bootstrap completion
  - User stories can then proceed in parallel
  - Or sequentially in priority order (`P1` → `P2` → `P3`)
- **Polish (Phase 7)**: Depends on all desired user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Runtime Bootstrap (Phase 3) - No dependencies on other stories
- **User Story 2 (P2)**: Can start after Runtime Bootstrap (Phase 3) - Builds on shared capture contracts but remains independently testable
- **User Story 3 (P3)**: Can start after Runtime Bootstrap (Phase 3) - Builds on shared audit and processing contracts but remains independently testable

### Within Each User Story

- Validation checklist and quickstart tasks come before implementation updates
- Contract and data-model updates come before repo-level summaries
- Repo-level summaries come before final cross-linking and end-to-end validation

### Parallel Opportunities

- Contract Cleanup tasks marked `[P]` can run in parallel
- Foundational tasks marked `[P]` can run in parallel
- Runtime Bootstrap tasks marked `[P]` can run in parallel by package
- Once Runtime Bootstrap completes, all three user stories can proceed in parallel if different people coordinate file ownership
- Validation tasks marked `[P]` within a story can run in parallel
- Contract, schema, and data-model updates marked `[P]` within a story can run in parallel before repo-level summary tasks

---

## Parallel Example: User Story 1

```bash
Task: "Add a checklist section for text capture durability and one-update-to-one-entry mapping in specs/002-telegram-capture/checklists/implementation.md"
Task: "Add a quickstart scenario for text capture confirmation and repository verification in specs/002-telegram-capture/quickstart.md"
Task: "Add a webhook smoke-validation scenario for text capture intake in specs/002-telegram-capture/quickstart.md"
```

```bash
Task: "Finalize daily inbox file rules, entry identity, and Obsidian readability expectations in specs/002-telegram-capture/contracts/telegram-capture-filesystem-contract.md"
Task: "Finalize accepted text input metadata, webhook intake semantics, and saved-status contract in specs/002-telegram-capture/contracts/telegram-capture-telegram-contract.md"
Task: "Document text message, daily inbox file, and capture buffer entry entities in specs/002-telegram-capture/data-model.md"
```

---

## Parallel Example: User Story 2

```bash
Task: "Add a checklist section for voice transcription confirmation, failure handling, and temporary-file cleanup in specs/002-telegram-capture/checklists/implementation.md"
Task: "Add a quickstart scenario for voice capture success, failure-aware status, and cleanup verification in specs/002-telegram-capture/quickstart.md"
Task: "Add a voice-processing audit review section covering transcription attempts and cleanup events in logs/action-log.md"
```

```bash
Task: "Finalize voice processing, user-visible confirmation, and cleanup rules in specs/002-telegram-capture/contracts/telegram-capture-processing-contract.md"
Task: "Document voice transcript, transcription attempt, and temporary voice artifact entities in specs/002-telegram-capture/data-model.md"
Task: "Define voice transcript, processing-status, and cleanup-report schema fields in schemas/workspace-config.schema.md"
```

---

## Parallel Example: User Story 3

```bash
Task: "Add a checklist section for burst traffic, repeated updates, retry-safe processing, repository write failure handling, and missing Stage 1 dependency behavior in specs/002-telegram-capture/checklists/implementation.md"
Task: "Add a quickstart scenario for rapid-message ordering, repeat visibility, retry validation, repository write failure behavior, and missing Stage 1 dependency behavior in specs/002-telegram-capture/quickstart.md"
Task: "Add an audit retention and repeat-activity review section in logs/action-log.md"
```

```bash
Task: "Finalize burst-traffic, repeat-update, and non-dedup posture in specs/002-telegram-capture/contracts/telegram-capture-telegram-contract.md"
Task: "Finalize retry semantics, failure preservation, webhook failure responses, and reprocessing boundaries in specs/002-telegram-capture/contracts/telegram-capture-processing-contract.md"
Task: "Document capture audit record, retry attempt, repeat-activity, and repository-write-failure entities in specs/002-telegram-capture/data-model.md"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Contract Cleanup
2. Complete Phase 2: Foundational Contracts
3. Complete Phase 3: Runtime Bootstrap
4. Complete Phase 4: User Story 1
5. Validate text capture durability with the checklist and quickstart artifacts
6. Stop and review before adding voice transcription complexity

### Incremental Delivery

1. Complete Contract Cleanup + Foundational Contracts to establish repo-level capture contracts
2. Complete Runtime Bootstrap so implementation targets concrete Go package paths
3. Deliver User Story 1 and validate daily inbox durability
4. Deliver User Story 2 and validate voice transcription plus cleanup behavior
5. Deliver User Story 3 and validate burst/retry auditability
6. Finish with cross-cutting consistency review

### Parallel Team Strategy

1. One person completes Contract Cleanup + Foundational Contracts + Runtime Bootstrap
2. After that:
   - Developer A: User Story 1
   - Developer B: User Story 2
   - Developer C: User Story 3
3. Stories complete independently and converge in Polish for final consistency checks
