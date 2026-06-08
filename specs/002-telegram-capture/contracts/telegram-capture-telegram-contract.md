# Contract: Telegram Capture Telegram Interface

## Purpose

Define the Telegram-facing expectations for Stage 2 input handling.

## Accepted Inputs

- text messages
- voice messages

## Required Metadata

- Telegram update identifier
- chat identifier
- message identifier
- sender identifier
- receive timestamp
- message kind
- optional raw text body for text inputs
- optional Telegram file identifier for voice inputs

## User-Visible Statuses

- text input accepted and saved
- voice input accepted, saved, and understood
- voice input saved but not fully understood due to transcription failure
- input accepted transport-wise but not durably saved to the workspace

## Burst and Repeat Handling

- Rapid consecutive messages are treated as separate inputs.
- Repeated Telegram updates are recorded for audit rather than silently dropped
  by semantic heuristics.
- Telegram remains the intake channel, but durable review stays possible through
  the repository.
- User-sent duplicate content is treated as a separate capture fact unless a
  later confirmed workflow merges it.

## Webhook Expectations

- Stage 2 intake uses a server-side webhook endpoint, not polling.
- Unsupported Telegram message kinds may be rejected with an audit-visible
  status, but they must not be misclassified as successfully captured text or
  voice input.
