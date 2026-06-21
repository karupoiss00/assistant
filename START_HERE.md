# Start Here

Минимальная цель первого этапа:

```text
Telegram -> OpenClaw on VPS -> append to Obsidian inbox -> git commit/push
```

## Что уже подготовлено в этом repo

- [docs/REMOTE_SETUP.md](/Users/rostislav.glizerin/projects/assistant/docs/REMOTE_SETUP.md)
- [deploy/scripts/bootstrap-vps.sh](/Users/rostislav.glizerin/projects/assistant/deploy/scripts/bootstrap-vps.sh)
- [deploy/templates/openclaw.json5.template](/Users/rostislav.glizerin/projects/assistant/deploy/templates/openclaw.json5.template)
- [deploy/templates/assistant.config.json.template](/Users/rostislav.glizerin/projects/assistant/deploy/templates/assistant.config.json.template)
- [deploy/templates/assistant.env.example](/Users/rostislav.glizerin/projects/assistant/deploy/templates/assistant.env.example)
- [deploy/templates/vault-bootstrap.sh.template](/Users/rostislav.glizerin/projects/assistant/deploy/templates/vault-bootstrap.sh.template)

## Нормальная схема репозиториев

1. Этот repo:
   контрольный repo для инфраструктуры и логики ассистента.
2. Отдельный `obsidian-vault` repo:
   только заметки и Markdown-структура.

Так действительно лучше, чем смешивать deploy-код и личный vault в одном месте.

## Первый запуск на VPS

1. Установить `Node 24`, `git`, `curl`.
2. Склонировать этот repo:

```bash
mkdir -p ~/assistant
git clone <this-repo-url> ~/assistant/control
cd ~/assistant/control
```

3. Запустить bootstrap:

```bash
bash deploy/scripts/bootstrap-vps.sh
```

4. Пройти официальный onboarding OpenClaw:

```bash
openclaw onboard --install-daemon
openclaw gateway status
openclaw doctor
```

5. Отредактировать:

```text
~/.openclaw/openclaw.json
~/assistant/config.json
~/assistant/.env
```

## Важные решения

- Gateway не публиковать наружу; держать `bind: "loopback"`.
- Telegram для одного владельца держать на `dmPolicy: "allowlist"`, а не на `open`.
- Vault пока не нужен для первого шага установки OpenClaw; его можно подключить вторым действием.

## Когда делать отдельный vault repo

Сразу после того, как:

1. OpenClaw уже установлен и жив как daemon.
2. Telegram бот отвечает только тебе.
3. Понятно, какой remote URL будет у vault repo.

Тогда на VPS:

```bash
cd ~/assistant
git clone git@github.com:USERNAME/obsidian-vault.git vault
bash ~/assistant/control/deploy/templates/vault-bootstrap.sh.template ~/assistant/vault
```
