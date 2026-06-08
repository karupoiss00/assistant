# Схема Конфигурации Workspace И Отчета

## Назначение

Задает формальную documentation schema для репозиторной конфигурации workspace и
payloads отчета нормализации.

## Поля Конфигурации Workspace

| Поле | Тип | Обязательно | Описание |
|-------|------|----------|-------------|
| `workspace_root_path` | string | да | Явный путь к production workspace repository |
| `log_retention_days` | integer | да | Минимальный срок хранения репозиторных логов; `>= 30` |
| `default_sphere_priority_order` | array<string> | да | Базовый порядок сфер для будущих ритуалов |
| `normalization_mode` | enum | да | `report-only` или `apply-safe-structure` |
| `confirmation_required_for_obsidian_changes` | boolean | да | Защита от переписывания или замены `.obsidian/` |
| `telegram_capture_enabled` | boolean | да | Включен ли Stage 2 Telegram capture runtime |
| `telegram_webhook_path` | string | да | HTTP path для Telegram webhook intake |
| `assistant_config_path` | string | да | Канонический путь к runtime-readable assistant config |
| `daily_inbox_path_pattern` | string | да | Канонический шаблон пути дневного inbox |
| `capture_audit_log_path_pattern` | string | да | Канонический шаблон пути capture audit log |
| `voice_transcription_required_for_confirmation` | boolean | да | Требуется ли транскрибация до user-visible success |
| `temporary_voice_cleanup_policy` | enum | да | Политика удаления временных voice artifacts |

## Правила Валидации Конфига

- `workspace_root_path` никогда не должен выводиться из `.workspace/`
- `log_retention_days` должен быть не меньше `30`
- `normalization_mode=apply-safe-structure` не разрешает destructive changes
- `default_sphere_priority_order` может включать больше сфер, чем набор `v1`,
  но не должен скрывать обязательные сферы
- `assistant_config_path` должен быть `System/assistant-config.yaml`
- `daily_inbox_path_pattern` должен разрешаться в
  `System/Inbox/YYYY/MM/YYYY-MM-DD.md`
- `capture_audit_log_path_pattern` должен разрешаться в
  `System/Logs/telegram-capture/YYYY/MM/YYYY-MM-DD.ndjson`
- `temporary_voice_cleanup_policy` для Stage 2 должна требовать удаления
  временного файла после success/failure

## Поля Отчета Нормализации

| Поле | Тип | Обязательно | Описание |
|-------|------|----------|-------------|
| `run_id` | string | да | Стабильный идентификатор отчета |
| `timestamp` | string | да | ISO-like timestamp запуска |
| `workspace_root_path` | string | да | Разрешенный путь к production workspace |
| `detected_violations` | array<object> | да | Отклонения от контракта, найденные при проверке |
| `executed_actions` | array<object> | да | Автоматически примененные структурные изменения |
| `proposed_actions` | array<object> | да | Предложения, не выполненные автоматически |
| `confirmation_gated_actions` | array<object> | да | Изменения, требующие явного подтверждения |
| `errors` | array<object> | да | Ошибки и причины пропуска путей |

## Ожидания К Полям Отчета

- Каждый action object содержит `action_id`, `target_path`, `autonomy_class`,
  `status` и `summary`
- Отчеты должны позволять различать структурные добавления и предложения,
  меняющие пользовательский смысл
- Отчеты должны поддерживать проверку planning-chain anchors и ожиданий по
  ручному редактированию через Obsidian, когда это релевантно

## Примеры

- Workspace, в котором отсутствует `Семья/Артефакты/`, может сначала появиться
  в `detected_violations`, а затем в `executed_actions`, если режим нормализации
  разрешает безопасное создание структуры.
- Предложение переименовать пользовательскую сферу появляется только в
  `proposed_actions`, если пользователь явно не подтвердил такое изменение.

## Поля Capture Audit Event

| Поле | Тип | Обязательно | Описание |
|-------|------|----------|-------------|
| `audit_id` | string | да | Стабильный идентификатор audit event |
| `event_timestamp` | string | да | Время события |
| `telegram_update_id` | string | да | Telegram update correlation key |
| `entry_id` | string | да | Durable inbox entry identifier |
| `attempt_id` | string | нет | Идентификатор processing attempt, если применимо |
| `event_type` | enum | да | `received`, `buffered`, `transcription_started`, `transcription_failed`, `transcription_succeeded`, `cleanup_completed`, `retry_requested`, `repeat_detected`, `workspace_write_failed` |
| `summary` | string | да | Человекочитаемое описание |

## Поля Daily Inbox Entry

| Поле | Тип | Обязательно | Описание |
|-------|------|----------|-------------|
| `entry_id` | string | да | Стабильный идентификатор inbox entry |
| `telegram_update_id` | string | да | Telegram update correlation key |
| `telegram_message_id` | string | да | Telegram message identifier |
| `message_kind` | enum | да | `text` или `voice` |
| `stored_text` | string | нет | Исходный текст для text capture |
| `transcript_text` | string | нет | Транскрипт для voice capture |
| `processing_status` | enum | да | `received`, `transcribing`, `understood`, `failed` |
| `receive_timestamp` | string | да | Время приема |
