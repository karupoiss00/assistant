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

## User-Visible Statuses

- text input accepted and saved
- voice input accepted and understood
- voice input saved but not fully understood due to transcription failure

## Burst and Repeat Handling

- Rapid consecutive messages are treated as separate inputs.
- Repeated Telegram updates are recorded for audit rather than silently dropped
  by semantic heuristics.
- Telegram remains the intake channel, but durable review stays possible through
  the repository.
