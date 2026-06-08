# Contract: Telegram Capture Filesystem

## Purpose

Define how Stage 2 stores Telegram capture state in the repository.

## Authoritative Repository State

- one daily inbox file for capture inputs
- repository-backed capture audit logs
- repository-managed capture configuration
- transcript text and processing status for voice inputs

## Production Workspace Layout

All Stage 2 durable writes happen inside the configured production workspace
root, not inside `.workspace/` of this repository.

| Artifact | Relative Path | Purpose | Authoritative |
|----------|---------------|---------|---------------|
| Assistant config | `System/assistant-config.yaml` | Runtime configuration and capture policy | yes |
| Daily inbox file | `System/Inbox/YYYY/MM/YYYY-MM-DD.md` | Append-only durable inbox for one calendar day | yes |
| Daily capture audit log | `System/Logs/telegram-capture/YYYY/MM/YYYY-MM-DD.ndjson` | Append-only event stream for capture processing and retries | yes |
| Temporary voice artifact | `runtime/<run_id>/<telegram_file_id>.ogg` | Ephemeral download used only during processing | no |

## Non-Authoritative State

- temporary voice files in `runtime/`
- transient network responses from Telegram or transcription services
- short-lived processing caches

## Daily Inbox Rules

- All inputs for the same day are appended to one daily inbox file.
- Each input remains individually identifiable inside the file.
- The inbox remains readable from Obsidian and is not replaced by Telegram as an
  exclusive interface.
- The inbox file path is derived only from the configured date-based path
  pattern, never from ad hoc runtime heuristics.
- Voice transcripts are stored inline in the inbox entry that represents the
  captured message; Stage 2 does not require a separate transcript file.

## Inbox Entry Rules

- Each inbox entry must have a stable `entry_id`.
- Each inbox entry must retain `telegram_update_id`, `message_id`, message kind,
  receive timestamp, and current processing status.
- Text inputs must store their accepted text directly in the entry.
- Voice inputs must store transcript text in the same entry once transcription
  succeeds.
- Retry or repeat activity must not overwrite the original entry identity.

## Audit Rules

- Audit records live outside the daily inbox file.
- Each inbox entry must be traceable to Telegram update identifiers and
  processing attempts.
- Audit retention must remain at least 30 days.
- Audit records are append-only and may reference an existing inbox entry
  through `entry_id`, `telegram_update_id`, and `attempt_id`.
- Repeated Telegram delivery and operator-initiated retry are separate audit
  facts and must not be collapsed into one event.
