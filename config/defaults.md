# Default Configuration

## Schedule
- Morning check-in: `08:00`
- Evening review: `23:00`

## Scope Priorities
- `Работа`
- `Здоровье`
- `Дом`
- `Семья`
- `Авто`

## Reviews
- Weekly review: enabled
- Monthly review: enabled

## Logging
- Retention: `30 days`
- Store logs in repository: yes

## Stage 2 Capture
- Telegram capture enabled: `true`
- Webhook path: `/telegram/webhook`
- Assistant config path: `System/assistant-config.yaml`
- Daily inbox path pattern: `System/Inbox/YYYY/MM/YYYY-MM-DD.md`
- Capture audit log path pattern: `System/Logs/telegram-capture/YYYY/MM/YYYY-MM-DD.ndjson`
- Voice transcription required for confirmation: `true`
- Temporary voice cleanup policy: `delete-after-processing`

## Notes
- Эти значения считаются стартовыми.
- Конкретный формат конфигурации будет зафиксирован отдельной технической спецификацией.
