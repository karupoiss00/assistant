# Implementation Plan: Telegram Capture

**Branch**: `002-name-telegram-capture` | **Date**: 2026-06-08 | **Spec**: [spec.md](./spec.md)

**Input**: Feature specification from `/specs/002-telegram-capture/spec.md`

## Summary

Build the first deployable Telegram entrypoint for the assistant so that text
and voice inputs are reliably captured into repository-backed durable state.
The implementation approach is repository-first: Telegram updates are received
through a bot, persisted into a daily inbox buffer plus audit logs, voice files
are transcribed before user-facing success confirmation, and temporary audio
artifacts are deleted after processing so that durable state remains compact and
Obsidian-compatible.

## Technical Context

**Language/Version**: Go for the Stage 2 webhook runtime plus Markdown contract
artifacts

**Primary Dependencies**: Go standard library HTTP/runtime primitives; Telegram
Bot API with webhook delivery; OpenAI speech-to-text API for voice
transcription; existing repository contracts in `config/`, `logs/`, `prompts/`,
and `schemas/`

**Storage**: File-based durable state in the repository; temporary local files
in `runtime/` only during voice processing

**Testing**: Integration-style validation for Telegram update intake,
repository-backed capture persistence, transcription success/failure handling,
and audit-log coverage

**Target Platform**: Long-running webhook service on a server or local operator
environment that can receive Telegram updates over HTTP and write to the
repository

**Project Type**: Stateful bot/service integrating Telegram, repository-backed
capture storage, and external transcription

**Performance Goals**: Text and voice capture should return user-visible
confirmation with basic understanding in under 10 seconds for 95% of validated
cases

**Constraints**: Durable state must remain in the repository; Telegram must not
become the only interface to stored inputs; voice confirmation is only complete
after transcription and basic understanding; temporary voice files must be
deleted after processing; no aggressive semantic deduplication in Stage 2

**Scale/Scope**: Single primary user or a very small set of trusted operators;
daily inbox files, capture logs, webhook intake, transcription flow,
retry-safe processing, and basic deployment-readiness for the first live
Telegram milestone

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
specs/002-telegram-capture/
├── plan.md
├── research.md
├── data-model.md
├── quickstart.md
├── contracts/
│   ├── telegram-capture-filesystem-contract.md
│   ├── telegram-capture-processing-contract.md
│   └── telegram-capture-telegram-contract.md
├── checklists/
│   ├── implementation.md
│   └── requirements.md
└── tasks.md
```

### Source Code (repository root)

```text
.workspace/                         # Reference workspace example only
cmd/                                # Go runtime entrypoints to be introduced during implementation
internal/                           # Go runtime packages to be introduced during implementation
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
runtime/                            # Temporary processing artifacts only
schemas/
├── README.md
└── workspace-config.schema.md
specs/
├── 001-workspace-foundation/
└── 002-telegram-capture/
AGENTS.md
SPEC.md
```

**Structure Decision**: Stage 2 still formalizes Telegram capture behavior
through explicit repository contracts, but it also requires a minimal live
runtime slice: a Go-based webhook service that receives updates, persists
durable state, triggers transcription, returns user-visible statuses, and keeps
temporary voice downloads in `runtime/` only.

## Phase 0: Research

### Research Goals

- Resolve the implementation posture for Telegram update delivery and repository
  persistence.
- Confirm webhook delivery as the minimum deployable runtime path.
- Use Go as the minimum runtime stack for the first live webhook service.
- Confirm the processing contract for voice transcription, user confirmation,
  and temporary-file cleanup.
- Clarify the technical runtime assumptions still missing from the repository.

### Research Output

- [research.md](./research.md) resolves delivery mode, capture storage shape,
  transcription lifecycle, and failure-handling strategy.

## Phase 1: Design & Contracts

### Planned Artifacts

- [data-model.md](./data-model.md): entities for Telegram messages, daily inbox
  files, transcription attempts, temporary voice artifacts, and audit records.
- [contracts/telegram-capture-filesystem-contract.md](./contracts/telegram-capture-filesystem-contract.md):
  repository paths and durable-state expectations for Stage 2 capture.
- [contracts/telegram-capture-processing-contract.md](./contracts/telegram-capture-processing-contract.md):
  processing flow, retries, cleanup rules, and user-visible confirmation
  semantics.
- [contracts/telegram-capture-telegram-contract.md](./contracts/telegram-capture-telegram-contract.md):
  Telegram-facing contract for update intake, metadata, and operator-visible
  statuses.
- [quickstart.md](./quickstart.md): validation guide for live Telegram tests,
  repository checks, transcription handling, and failure scenarios.
- [checklists/implementation.md](./checklists/implementation.md): execution-time
  validation checklist for contract coverage, webhook smoke checks, and final
  Stage 2 walkthroughs.

### Planned Mutations

- `auto-allowed`: Save each Telegram event into the daily inbox file, append
  audit-log records, save transcription text, mark processing attempts, and
  delete temporary voice artifacts after completion or failure.
- `proposal-only`: Suggest semantic classification, redistribution into richer
  workspace artifacts, or later deduplication of similar inputs.
- `requires-confirmation`: Delete durable inbox entries, merge distinct inputs
  into one semantic object, rewrite already stored user meaning, or introduce
  long-lived storage of original voice artifacts as product policy.

### Ritual Coverage

- This feature improves the daily capture ritual by making Telegram a reliable
  ingestion channel into the repository.
- It preserves cross-sphere balance by storing inputs before classification
  rather than biasing them toward one sphere too early.

### Logging and Audit

- Every Telegram update, processing attempt, transcription outcome, retry, and
  cleanup action must be recorded in repository-backed logs.
- Minimum retention remains 30 days.
- Operators must be able to trace a Telegram input from update identifiers to
  inbox content, transcription outcome, and cleanup status.

### Required Contract Updates Before Implementation

- `config/`: define Telegram capture configuration expectations such as bot
  enablement, inbox path conventions, and confirmation behavior.
- `logs/`: define Stage 2 capture event types and correlations to Telegram
  update IDs.
- `schemas/`: define daily inbox entry and capture log/report schemas.
- `prompts/`: add or update prompt contracts for downstream classification of
  captured content, without making classification part of Stage 2 completion.
- Workspace note formats: define how daily inbox files fit into the repository
  contract without breaking Stage 1 source-of-truth boundaries.
- Webhook runtime contract: define the minimum live intake path from Telegram
  webhook to repository mutation, user reply, and temporary-file cleanup.

### Validation Strategy

- Validate that text inputs land in the daily inbox file and audit logs.
- Validate that voice inputs produce user-facing confirmation only after
  transcription and basic understanding.
- Validate that transcription failure still preserves durable state and logs the
  failure.
- Validate that temporary voice files are removed after processing.
- Validate that rapid message bursts do not lose ordering or content.
- Validate that a live Telegram webhook request can reach the service and
  produce the expected repository mutation plus user-visible reply.

## Post-Design Constitution Check

- [x] Authoritative repository files and directories are identified in the plan,
      data model, and capture contracts.
- [x] Planned mutations are classified and remain aligned with repository-first
      durability plus confirmed autonomy.
- [x] Ritual coverage is explicit: Stage 2 exists to make daily capture
      reliable.
- [x] Logging, retention, and traceability are defined for both success and
      failure paths.
- [x] Contract updates for config, prompts, schemas, logs, and note formats are
      listed before implementation.
- [x] Validation checks user-visible confirmation behavior and repository
      mutations together.

## Complexity Tracking

No constitution violations require justification for this stage.
