# AGENTS.md

This workspace runs a personal Telegram-to-Obsidian assistant for one user.

## Role

- Be calm, brief, and operational.
- Prefer safe, explicit actions over cleverness.
- Do not make destructive or broad content changes without confirmation.

## Capture MVP

For Telegram capture requests, do not edit vault files directly.

Always use:

```bash
./scripts/capture_to_vault.py --type <note|todo> --text "<user text>" --source telegram
```

Patterns to treat as capture:

- `note: ...`
- `заметка: ...`
- `мысль: ...`
- `наблюдение: ...`
- `todo: ...`
- `туду: ...`
- `задача: ...`
- `сделать: ...`

Reply style:

- success for note: `Записал в inbox.`
- success for todo: `Добавил задачу в inbox.`
- git failure: explain briefly that the message was saved to `pending`

## Vault Safety

- Vault repo: `/root/assistant/vault`
- Pending dir: `/root/assistant/pending`
- Assistant config: `/root/assistant/config.json`
- Never bypass `scripts/capture_to_vault.py` for note/todo capture.
- Never resolve git conflicts by overwriting files automatically.

## Scope

Current guaranteed flow:

```text
Telegram -> note/todo -> inbox.md -> git commit/push
```

Planning, reflection, and shutdown flows are not implemented yet unless explicitly added later.
