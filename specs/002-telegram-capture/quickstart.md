# Quickstart: Telegram Capture Validation

## Purpose

Use this guide to validate that Stage 2 delivers a deployable Telegram capture
loop and preserves repository-backed durable state.

## Prerequisites

- Feature branch `002-name-telegram-capture`
- Access to the capture contracts in [contracts/](./contracts/)
- Review access to [data-model.md](./data-model.md) and [plan.md](./plan.md)
- Review access to [checklists/implementation.md](./checklists/implementation.md)
- A configured Telegram bot token and OpenAI transcription access for live
  validation once implementation exists

## Scenario 1: Validate text capture durability

1. Open [contracts/telegram-capture-filesystem-contract.md](./contracts/telegram-capture-filesystem-contract.md).
2. Confirm the contract defines `System/Inbox/YYYY/MM/YYYY-MM-DD.md` as the
   daily inbox path and a separate daily audit log path.
3. Review [data-model.md](./data-model.md) and confirm that one text Telegram
   update maps to one inbox entry and one audit trail.
4. Open [checklists/implementation.md](./checklists/implementation.md) and
   confirm the text durability checks match the contract.
5. Validate that repository-backed state, not runtime state, is authoritative.

**Expected outcome**:
- Text capture lands in a daily inbox file.
- Audit records remain separate and traceable.
- User-visible success depends on durable writes, not transport-only receipt.

## Scenario 1a: Validate text webhook smoke behavior

1. Confirm [plan.md](./plan.md) names a server-side webhook service as the
   Stage 2 ingress path.
2. Confirm [contracts/telegram-capture-telegram-contract.md](./contracts/telegram-capture-telegram-contract.md)
   explicitly requires webhook intake rather than polling.
3. Confirm [tasks.md](./tasks.md) routes text intake implementation through
   `internal/telegram/` and `internal/capture/`.

**Expected outcome**:
- A reviewer can identify the concrete webhook ingress path before runtime code
  exists.

## Scenario 1b: Validate text latency target coverage

1. Review [spec.md](./spec.md) and confirm the 10-second target remains part of
   Stage 2 success criteria.
2. Review [checklists/implementation.md](./checklists/implementation.md) and
   confirm text success depends on durable save rather than early transport
   acknowledgment.
3. Review [tasks.md](./tasks.md) and confirm there is an explicit task to record
   latency results against the 95% / 10-second target.

**Expected outcome**:
- Text latency validation has an explicit target and an implementation surface.

## Scenario 2: Validate voice capture confirmation behavior

1. Open [contracts/telegram-capture-processing-contract.md](./contracts/telegram-capture-processing-contract.md).
2. Confirm that user-visible success for voice capture happens only after
   transcription and basic understanding.
3. Confirm that the original voice file is temporary and must be deleted after
   processing.
4. Review [data-model.md](./data-model.md) and confirm the transcription
   attempt state model matches the contract.
5. Confirm the checklist requires `cleanup_completed` audit visibility.

**Expected outcome**:
- Voice capture is not considered successfully acknowledged before transcript
  availability.
- Temporary voice files are excluded from durable repository state.

## Scenario 2a: Validate voice webhook smoke behavior

1. Confirm [contracts/telegram-capture-telegram-contract.md](./contracts/telegram-capture-telegram-contract.md)
   lists voice messages as accepted webhook inputs.
2. Confirm [contracts/telegram-capture-processing-contract.md](./contracts/telegram-capture-processing-contract.md)
   sequences inbox write, transcription, reply, and cleanup.
3. Confirm [tasks.md](./tasks.md) routes voice work through `internal/capture/`,
   `internal/transcribe/`, and `internal/workspace/`.

**Expected outcome**:
- Voice capture implementation has a concrete intake and processing path.

## Scenario 2b: Validate voice latency target coverage

1. Review [spec.md](./spec.md) and confirm the user-visible 10-second target
   also applies to voice confirmation.
2. Confirm the processing contract requires transcription before final success.
3. Confirm the implementation checklist makes transcript persistence and cleanup
   part of completion.
4. Confirm the Stage 2 definition of `basic understanding` is narrow and
   testable.

**Expected outcome**:
- Voice latency is treated as a measured contract, not a best-effort note.
- Voice success semantics are testable without requiring full semantic routing.

## Scenario 3: Validate failure and retry safety

1. Open [contracts/telegram-capture-processing-contract.md](./contracts/telegram-capture-processing-contract.md).
2. Walk through a failed transcription path.
3. Confirm the inbox entry and audit trail survive the failure.
4. Confirm retry handling reuses the saved durable context rather than requiring
   the original temporary voice artifact to stay forever.
5. Confirm user-visible failure wording distinguishes "saved" from
   "understood".

**Expected outcome**:
- Failures do not lose the captured input.
- Operators can retry based on durable context and logs.

## Scenario 4: Validate burst traffic and duplicate posture

1. Open [contracts/telegram-capture-telegram-contract.md](./contracts/telegram-capture-telegram-contract.md).
2. Confirm rapid consecutive messages are stored as distinct capture facts.
3. Confirm repeated Telegram delivery is logged without aggressive semantic
   deduplication.
4. Confirm workspace write failures are represented separately from repeat
   activity.

**Expected outcome**:
- Burst traffic does not collapse inputs into one record.
- Deduplication remains deferred and observable.

## Scenario 4a: Validate retry and workspace failure behavior

1. Open [contracts/telegram-capture-processing-contract.md](./contracts/telegram-capture-processing-contract.md).
2. Confirm retry creates a new `attempt_id` while preserving the original
   durable entry.
3. Confirm workspace write failure is represented separately from successful
   capture and from repeat delivery.
4. Confirm [logs/action-log.md](../../logs/action-log.md) includes
   `workspace_write_failed` and retry-related audit expectations.

**Expected outcome**:
- Retry, repeat, and workspace failure paths are distinguishable in both
  contracts and audit semantics.

## Scenario 5: Validate runtime bootstrap scope

1. Open [plan.md](./plan.md) and confirm Stage 2 uses a long-running Go webhook
   service.
2. Open [tasks.md](./tasks.md) and confirm runtime work names concrete source
   paths under `cmd/` and `internal/`.
3. Confirm the implementation checklist includes runtime bootstrap readiness.

**Expected outcome**:
- The implementation kickoff can start from concrete code paths instead of an
  abstract "runtime service".
