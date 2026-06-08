# Contract: Telegram Capture Filesystem

## Purpose

Define how Stage 2 stores Telegram capture state in the repository.

## Authoritative Repository State

- one daily inbox file for capture inputs
- repository-backed capture audit logs
- repository-managed capture configuration
- transcript text and processing status for voice inputs

## Non-Authoritative State

- temporary voice files in `runtime/`
- transient network responses from Telegram or transcription services
- short-lived processing caches

## Daily Inbox Rules

- All inputs for the same day are appended to one daily inbox file.
- Each input remains individually identifiable inside the file.
- The inbox remains readable from Obsidian and is not replaced by Telegram as an
  exclusive interface.

## Audit Rules

- Audit records live outside the daily inbox file.
- Each inbox entry must be traceable to Telegram update identifiers and
  processing attempts.
- Audit retention must remain at least 30 days.
