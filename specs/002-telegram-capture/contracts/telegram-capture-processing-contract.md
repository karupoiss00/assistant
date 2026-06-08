# Contract: Telegram Capture Processing

## Purpose

Define the Stage 2 processing flow for text and voice capture.

## Text Processing

- Receive Telegram text input
- Persist it into the daily inbox file
- Create audit records
- Return user-visible confirmation after durable save

## Voice Processing

- Receive Telegram voice input metadata
- Download a temporary voice artifact to `runtime/`
- Start transcription
- Persist transcript, metadata, and audit information
- Return user-visible success only after transcript and basic understanding are
  available
- Delete the temporary voice artifact after success or failure handling

## Failure Handling

- If transcription fails, the durable inbox record remains.
- Failure is written to audit logs.
- The user receives a failure-aware status explaining that the message was not
  fully understood yet.
- Retry uses the durable inbox context and a new processing attempt.

## Deduplication Posture

- Stage 2 does not perform aggressive semantic deduplication.
- Repeated Telegram delivery may be logged as repeat activity.
- Inputs are not automatically merged based on heuristic similarity.
