# Remote Setup

Этот репозиторий используется как control repo для деплоя и конфигурации OpenClaw на VPS.

Отдельно:

- этот repo: bootstrap-скрипты, шаблоны конфигов, инструкции, prompts/skills для ассистента;
- будущий `obsidian-vault` repo: только Markdown vault, который будет клонироваться на VPS в `~/assistant/vault`.

## Что хранить здесь

- `deploy/scripts/` - shell-скрипты для первичной настройки VPS;
- `deploy/templates/` - шаблоны OpenClaw и assistant config;
- `prompts/` - системные промпты и runbook-заметки для ассистента;
- `skills/` - workspace skills для OpenClaw;
- `docs/` - инструкции по запуску, обновлению и отладке.

## Что не хранить здесь

- содержимое Obsidian vault;
- `.env` с боевыми секретами;
- SSH private keys;
- runtime state OpenClaw из `~/.openclaw/`.

## Рекомендуемая структура на VPS

```text
~/assistant/
  control/                  # clone этого репо
  vault/                    # clone отдельного obsidian-vault repo
  pending/
  logs/
  scripts/
  prompts/
  skills/
  config.json
  .env
```

## Актуальные допущения по OpenClaw

По официальному README и docs на 2026-06-22:

- рекомендуемый runtime: `Node 24`, допустимо `Node 22.19+`;
- рекомендуемая установка: `npm install -g openclaw@latest`;
- рекомендуемый первый запуск: `openclaw onboard --install-daemon`;
- основной config OpenClaw хранится в `~/.openclaw/openclaw.json`;
- Telegram для single-user бота лучше настраивать через `dmPolicy: "allowlist"` и `allowFrom: ["<numeric-user-id>"]`;
- встроенный scheduler OpenClaw лучше обычного system cron для agent-задач.

## Первые практические шаги

1. Поднять VPS и завести пользователя без root-работы по умолчанию.
2. Склонировать этот repo в `~/assistant/control`.
3. Выполнить `deploy/scripts/bootstrap-vps.sh`.
4. Пройти `openclaw onboard --install-daemon`, но не открывать gateway наружу.
5. Настроить Telegram bot token и allowlist.
6. Позже создать отдельный `obsidian-vault` repo и склонировать его в `~/assistant/vault`.

## Почему vault отдельным repo

- проще изолировать личные заметки от deploy-кода;
- проще дать deploy key только на один vault repo;
- меньше риск случайно закоммитить секреты или runtime-файлы OpenClaw;
- можно менять логику ассистента независимо от содержимого заметок.
