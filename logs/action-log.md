# Контракт Журнала Действий

## Назначение

Определяет репозиторный audit-контракт для нормализации workspace и связанных
действий ассистента.

## Обязательные категории логов

- нарушения контракта, найденные при проверке workspace
- автоматически примененные структурные исправления
- предложенные, но не выполненные изменения
- действия, требующие подтверждения, и их итоговое решение
- ошибки нормализации и пропущенные пути
- Stage 2 capture events: `received`, `buffered`, `transcription_started`,
  `transcription_failed`, `transcription_succeeded`, `cleanup_completed`,
  `retry_requested`, `repeat_detected`, `workspace_write_failed`

## Обязательные поля

- `log_entry_id`: стабильный идентификатор
- `timestamp`: время события
- `workspace_root_path`: разрешенный путь к production workspace
- `run_id`: идентификатор запуска нормализации/отчета
- `event_type`: `contract_violation`, `auto_fix`, `proposal`,
  `confirmed_change`, `error`
- `action_ids`: идентификаторы связанных действий нормализации
- `target_paths`: затронутые пути в репозитории
- `summary`: человекочитаемое описание
- `autonomy_class`: `auto-allowed`, `proposal-only`,
  `requires-confirmation`
- `status`: `executed`, `proposed`, `confirmed`, `rejected`, `failed`
- `telegram_update_id`: корреляция с Telegram intake, если событие относится к
  capture-слою
- `entry_id`: идентификатор durable inbox entry, если применимо
- `attempt_id`: идентификатор processing attempt, если применимо

## Чеклист Проверки Отчета Нормализации

- [x] Отчет фиксирует разрешенный путь к production workspace
- [x] Отчет различает автоматически выполненные, предложенные и подтвержденные действия
- [x] Отчет перечисляет target paths для каждой мутации или пропущенного пути
- [x] Отчет фиксирует ошибки и причины пропуска работы
- [x] Отчет можно сопоставить с изменениями в репозитории и Git history

## Хранение

- Логи должны храниться в репозиторном хранилище не менее 30 дней.
- Контракт логов должен оставаться проверяемым без чтения runtime-кода.
- Git history — это дополнительное audit-доказательство, а не замена
  семантическим log records.
- Для Stage 2 capture append-only audit log в production workspace считается
  канонической операционной записью, а этот файл фиксирует его contract shape.

## Ожидания Оператора

- Пользователь может понять из журнала, что изменилось и почему.
- Оператор может отследить отчет до конкретных затронутых путей.
- Логи должны оставаться совместимыми с будущей формализацией схем в `schemas/`.
