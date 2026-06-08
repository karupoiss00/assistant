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

## Правила Валидации Конфига

- `workspace_root_path` никогда не должен выводиться из `.workspace/`
- `log_retention_days` должен быть не меньше `30`
- `normalization_mode=apply-safe-structure` не разрешает destructive changes
- `default_sphere_priority_order` может включать больше сфер, чем набор `v1`,
  но не должен скрывать обязательные сферы

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
