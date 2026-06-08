# Quickstart: Telegram Capture Validation

## Purpose

Use this guide to validate that Stage 2 delivers a deployable Telegram capture
loop and preserves repository-backed durable state.

## Prerequisites

- Feature branch `002-name-telegram-capture`
- Access to the capture contracts in [contracts/](./contracts/)
- Review access to [data-model.md](./data-model.md) and [plan.md](./plan.md)
- A configured Telegram bot token and OpenAI transcription access for live
  validation once implementation exists

## Scenario 1: Validate text capture durability

1. Open [contracts/telegram-capture-filesystem-contract.md](./contracts/telegram-capture-filesystem-contract.md).
2. Confirm the contract defines one daily inbox file plus separate audit logs.
3. Review [data-model.md](./data-model.md) and confirm that one text Telegram
   update maps to one inbox entry and one audit trail.
4. Validate that repository-backed state, not runtime state, is authoritative.

**Expected outcome**:
- Text capture lands in a daily inbox file.
- Audit records remain separate and traceable.

## Scenario 2: Validate voice capture confirmation behavior

1. Open [contracts/telegram-capture-processing-contract.md](./contracts/telegram-capture-processing-contract.md).
2. Confirm that user-visible success for voice capture happens only after
   transcription and basic understanding.
3. Confirm that the original voice file is temporary and must be deleted after
   processing.
4. Review [data-model.md](./data-model.md) and confirm the transcription
   attempt state model matches the contract.

**Expected outcome**:
- Voice capture is not considered successfully acknowledged before transcript
  availability.
- Temporary voice files are excluded from durable repository state.

## Scenario 3: Validate failure and retry safety

1. Open [contracts/telegram-capture-processing-contract.md](./contracts/telegram-capture-processing-contract.md).
2. Walk through a failed transcription path.
3. Confirm the inbox entry and audit trail survive the failure.
4. Confirm retry handling reuses the saved durable context rather than requiring
   the original temporary voice artifact to stay forever.

**Expected outcome**:
- Failures do not lose the captured input.
- Operators can retry based on durable context and logs.

## Scenario 4: Validate burst traffic and duplicate posture

1. Open [contracts/telegram-capture-telegram-contract.md](./contracts/telegram-capture-telegram-contract.md).
2. Confirm rapid consecutive messages are stored as distinct capture facts.
3. Confirm repeated Telegram delivery is logged without aggressive semantic
   deduplication.

**Expected outcome**:
- Burst traffic does not collapse inputs into one record.
- Deduplication remains deferred and observable.
