# Implementation Checklist: Telegram Capture

## Purpose

Use this checklist during implementation and final validation of Stage 2.

## Text Capture Durability

- [ ] One accepted text update produces exactly one inbox entry.
- [ ] The text inbox entry is written to `System/Inbox/YYYY/MM/YYYY-MM-DD.md`.
- [ ] At least `received` and `buffered` audit events are appended for the text
      update.
- [ ] User-visible success for text is sent only after durable inbox and audit
      writes succeed.

## Voice Capture Success

- [ ] One accepted voice update produces exactly one inbox entry before
      transcription starts.
- [ ] The same inbox entry is updated with transcript text after successful
      transcription.
- [ ] Audit events include `received`, `buffered`, `transcription_started`,
      `transcription_succeeded`, and `cleanup_completed`.
- [ ] User-visible success for voice is sent only after transcript persistence
      and audit writes succeed.

## Voice Failure Handling

- [ ] Failed transcription leaves the inbox entry durable and traceable.
- [ ] Failure path appends `transcription_failed` audit event.
- [ ] Failure-aware user reply states that the message was saved but not fully
      understood.
- [ ] Retry uses the saved entry context and generates a new `attempt_id`.

## Temporary File Cleanup

- [ ] Temporary voice artifacts are created only under `runtime/`.
- [ ] Temporary voice artifacts are deleted after both success and failure
      paths.
- [ ] Temporary voice artifacts never become authoritative repository state.

## Burst And Repeat Handling

- [ ] A burst of at least 5 rapid messages produces 5 distinct inbox entries.
- [ ] Entry order in the daily inbox matches accepted message order.
- [ ] Repeated Telegram delivery is audit-visible as repeat activity.
- [ ] User-sent duplicate content is preserved as separate capture facts.

## Workspace Failure Paths

- [ ] If the workspace write fails, the user does not receive a false success.
- [ ] If required Stage 1 paths are missing, the failure is audit-visible.
- [ ] Partial processing states are distinguishable from successful durable
      capture.

## Runtime Bootstrap Readiness

- [ ] `go.mod` exists for the Stage 2 runtime.
- [ ] `cmd/telegram-capture/` contains the runtime entrypoint.
- [ ] Internal package boundaries exist for config, Telegram intake, capture
      writes, audit logging, transcription, and workspace access.
