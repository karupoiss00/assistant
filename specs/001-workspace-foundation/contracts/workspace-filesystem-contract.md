# Contract: Workspace Filesystem

## Purpose

Define the authoritative `v1` filesystem contract for an Obsidian-compatible
workspace repository used by the assistant.

## Root Rules

- The production workspace is an Obsidian-compatible Git repository whose root
  path is configured externally to this document.
- `.workspace/` in this repository is a reference example only.
- Durable user state lives in the production workspace repository, not in
  `runtime/`.

## Required `v1` Sphere Directories

The workspace must support these top-level spheres:

- `Работа`
- `Дом`
- `Авто`
- `Здоровье`
- `Семья`

Each required sphere must contain:

- `Цели.md`
- `Тактические цели.md`
- `Операционные задачи.md`
- `Артефакты/`

## Required Shared Artifacts

- `Цели на 5 лет.md`
- Shared sphere-template reference matching the semantics of
  `.workspace/!Общее/Шаблон сферы`
- A repository-managed assistant configuration location
- A repository-managed assistant action-log location
- `.obsidian/` as preserved client state when present

## Template Semantics

The reference template defines expected note roles:

- `Цели.md`: long-horizon sphere goals with session date and target date
- `Тактические цели.md`: current-month tactical goals with one key monthly task
- `Операционные задачи.md`: operational backlog items that may link to notes in
  `Артефакты/`

The exact production path for the template may differ from
`.workspace/!Общее/Шаблон сферы`, but the semantic contract must remain
equivalent.

## Link Integrity Expectations

The contract must preserve explicit anchor points for:

- operational task
- monthly goal
- annual or long-horizon sphere goal
- `Цель на 5 лет`

Later stages may refine the concrete link syntax, but normalization may not
silently destroy existing note links, attachment references, or artifact paths.

The contract also treats the chain
`Операционная задача -> Цель месяца -> Годовая цель -> Цель на 5 лет` as a
required planning-horizon backbone for `v1`, even before later stages define
the full linking mechanics.

## Manual Obsidian Readability

- Required artifacts must remain readable and editable directly through
  Obsidian.
- No required contract artifact may depend on Telegram as the only editing
  interface.
- Structural normalization may add missing elements, but it may not make the
  resulting note set harder to inspect manually.

## Non-Required Existing Content

- Additional top-level spheres are allowed.
- Additional notes, attachments, and directories inside required spheres are
  allowed.
- Existing user content outside the required contract remains authoritative and
  must not be deleted merely because it is outside `v1` minimum structure.
