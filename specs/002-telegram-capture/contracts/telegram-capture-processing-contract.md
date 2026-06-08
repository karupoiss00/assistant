# Contract: Telegram Capture Processing

## Purpose

Define the Stage 2 processing flow for text and voice capture.

## Text Processing

1. Receive Telegram text input through the webhook endpoint.
2. Validate that the update contains a supported text payload.
3. Append a new inbox entry to the daily inbox file.
4. Append at least `received` and `buffered` audit events.
5. Return user-visible confirmation only after the inbox write and audit write
   succeed.

## Voice Processing

1. Receive Telegram voice input metadata through the webhook endpoint.
2. Append a new inbox entry with `processing_status=received`.
3. Append `received` and `buffered` audit events.
4. Download a temporary voice artifact to `runtime/`.
5. Append `transcription_started` audit event and set
   `processing_status=transcribing`.
6. Call the transcription dependency.
7. On success, persist transcript text and `processing_status=understood` into
   the same inbox entry and append `transcription_succeeded`.
8. Return user-visible success only after transcript persistence, audit write,
   and basic understanding are available.
9. Delete the temporary voice artifact and append `cleanup_completed`.

## Failure Handling

- If transcription fails, the durable inbox record remains and the original
  `entry_id` is preserved.
- Failure is written to audit logs as `transcription_failed`.
- The inbox entry remains available for retry with `processing_status=failed`.
- The user receives a failure-aware status explaining that the message was
  saved but not fully understood.
- Retry uses the durable inbox context and a new `attempt_id`.
- If durable write to the workspace fails, the user must not receive a success
  status even if Telegram delivery or transcription already succeeded.
- If required Stage 1 workspace artifacts or configured paths are missing, the
  runtime records an error path and reports partial processing to the user.

## User-Visible Status Semantics

- `transport accepted`: Telegram delivered the request to the webhook, but this
  is not yet product success.
- `durable saved`: text input or pre-transcription voice entry was saved to the
  workspace.
- `understood`: transcript and durable state are both available and the bot can
  acknowledge the content meaningfully.
- `saved but not understood`: durable state exists, but transcription or basic
  understanding failed.
- `not durably saved`: the transport path may have succeeded, but the workspace
  was not updated correctly.

## Definition Of Basic Understanding

- Stage 2 does not require full semantic classification to mark a voice input as
  understood.
- `basic understanding` is reached only when:
  - transcript text is non-empty;
  - transcript text is durably persisted into the inbox entry;
  - the user-facing reply can acknowledge that the message was captured and
    transcribed using a deterministic confirmation format.
- A transcript that is empty, missing, or not durably written cannot be treated
  as understood.

## Deduplication Posture

- Stage 2 does not perform aggressive semantic deduplication.
- Repeated Telegram delivery may be logged as repeat activity.
- Inputs are not automatically merged based on heuristic similarity.
