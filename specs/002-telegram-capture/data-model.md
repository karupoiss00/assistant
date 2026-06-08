# Data Model: Telegram Capture

## Telegram Incoming Message

- **Description**: A Telegram-delivered input event representing either text or
  voice capture.
- **Fields**:
  - `telegram_update_id`
  - `chat_id`
  - `message_id`
  - `sender_id`
  - `received_at`
  - `message_kind`: `text` or `voice`
  - `raw_text`: optional text body from Telegram
  - `voice_file_reference`: optional Telegram file reference
  - `processing_status`: `received`, `transcribing`, `understood`, `failed`
- **Validation Rules**:
  - `telegram_update_id` and `message_id` must be retained for audit
  - Every incoming message must map to one daily inbox entry

## Daily Inbox File

- **Description**: Repository-backed daily buffer where all capture inputs for a
  calendar day are appended.
- **Fields**:
  - `date_key`
  - `entries`: ordered list of capture buffer entries
  - `entry_count`
- **Validation Rules**:
  - One active daily inbox file per day
  - Entries must retain order of accepted capture events
  - Daily inbox content remains readable from Obsidian

## Capture Buffer Entry

- **Description**: One stored capture item inside the daily inbox file.
- **Fields**:
  - `entry_id`
  - `telegram_update_id`
  - `message_kind`
  - `stored_text`
  - `transcript_text`: optional
  - `metadata_ref`
  - `processing_status`
  - `created_at`
- **Validation Rules**:
  - `stored_text` must exist for text inputs
  - Voice inputs become user-visible success only after `transcript_text` exists
    and basic understanding is available

## Voice Transcription Attempt

- **Description**: One attempt to convert a temporary voice artifact into text.
- **Fields**:
  - `attempt_id`
  - `entry_id`
  - `started_at`
  - `finished_at`
  - `outcome`: `success`, `failure`
  - `error_summary`: optional
  - `transcript_text`: optional
- **State Transitions**:
  - `received -> transcribing -> understood`
  - `received -> transcribing -> failed`
- **Validation Rules**:
  - Failure must not delete the durable capture record
  - Success must produce transcript text before user-visible success confirmation

## Temporary Voice Artifact

- **Description**: Ephemeral downloaded voice file used only during processing.
- **Fields**:
  - `artifact_id`
  - `telegram_file_id`
  - `local_runtime_path`
  - `created_at`
  - `deleted_at`
- **Validation Rules**:
  - Temporary artifacts live only in `runtime/`
  - Temporary artifacts must be deleted after success or failure handling
  - Temporary artifacts are never part of authoritative repository state

## Capture Audit Record

- **Description**: Repository-backed event record tied to Telegram capture.
- **Fields**:
  - `audit_id`
  - `telegram_update_id`
  - `entry_id`
  - `event_type`: `received`, `buffered`, `transcription_started`,
    `transcription_failed`, `transcription_succeeded`, `cleanup_completed`,
    `retry_requested`
  - `event_timestamp`
  - `summary`
- **Validation Rules**:
  - Every input must have at least `received` and `buffered` audit events
  - Voice inputs must also log transcription and cleanup outcomes

## Capture Configuration

- **Description**: Repository-managed configuration governing Telegram capture.
- **Fields**:
  - `telegram_capture_enabled`
  - `daily_inbox_path_pattern`
  - `log_retention_days`
  - `voice_transcription_required_for_confirmation`
  - `temporary_voice_cleanup_policy`
- **Validation Rules**:
  - `log_retention_days >= 30`
  - `voice_transcription_required_for_confirmation` is `true` for Stage 2
  - `temporary_voice_cleanup_policy` must define deletion after processing
