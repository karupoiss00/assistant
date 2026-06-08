# Open Questions: Stage 2 and Future Runtime

## Зачем этот файл

Этот файл фиксирует архитектурные вопросы и ожидания, которые возникли во
время Stage 2, но еще не стали принятыми контрактами реализации.

## Что уже понятно сейчас

- После Stage 2 ожидается первый живой Telegram milestone: бот принимает текст
  и голос, сохраняет входящее в репозиторный контур и не теряет сообщения.
- Durable state должен жить в Obsidian-совместимом git-репозитории, а не в
  runtime-памяти сервиса.
- Сам Stage 2 пока фиксирует в первую очередь capture-контур, а не полную
  оркестрацию по сферам, ритуалам и смысловым связям.

## Решение 1: Что будет "в конце" ближайшего этапа

Зафиксированное решение:

- Stage 2 заканчивается минимальным deployable long-running webhook-сервисом на
  сервере/VPS, подключенным к Telegram Bot API.
- Этот сервис отвечает только за capture-контур: прием апдейтов, durable
  запись в production workspace, транскрибацию голоса, user-visible reply и
  audit trail.
- Stage 2 не включает автоматизацию morning/evening rituals, weekly/monthly
  review или полную orchestration-логику.

Что остается за пределами текущего решения:

- конкретный deployment recipe;
- supervisor/restart policy;
- operational lifecycle и наблюдаемость production-процесса.

## Вопрос 2: Где будет жить рабочее Obsidian-пространство

Текущее понимание:

- Production workspace не обязан жить внутри этого репозитория.
- По Stage 1 workspace должен быть отдельным git-репозиторием с
  конфигурируемым путем.
- `.workspace/` в текущем проекте остается reference example, а не production
  source of truth.

Рабочая гипотеза для будущей реализации:

- на сервере/VPS будет лежать:
  - этот репозиторий ассистента;
  - отдельно пользовательский Obsidian/workspace-репозиторий;
- ассистент будет работать с workspace по явно заданному пути из конфигурации.

## Решение 3: Будет ли Codex "развернут на сервере"

Зафиксированное решение:

- Stage 2 не требует постоянного Codex-runner как production runtime.
- Минимальный production runtime Stage 2 - обычный long-running Go webhook
  service, который при необходимости вызывает внешние API.
- Codex остается допустимым операторским и разработческим контуром, но не
  является обязательной частью runtime topology этого этапа.

## Решение 4: Сможет ли ассистент управлять workspace через skill и Obsidian CLI

Зафиксированное решение:

- Stage 2 runtime пишет в production workspace напрямую как файловый writer в
  рамках явных repository contracts.
- Stage 2 не зависит от Obsidian CLI, специальных skills или отдельного
  orchestration-слоя для выполнения capture-записей.
- Более сложная модель управления workspace через skills, CLI или смешанный
  automation-контур остается темой Stage 3-4.

## Вопрос 5: На каком этапе остальные runtime-решения должны быть приняты

Предварительное ожидание:

- Stage 2: минимальный deployable Telegram capture runtime;
- Stage 3: появляется ежедневный рабочий цикл и становится важнее точная модель
  фонового процесса и планировщика;
- Stage 4-5: становится критично, как агент читает, обновляет и оркестрирует
  workspace;
- до Stage 4 желательно явно зафиксировать runtime topology, deployment model и
  правила изменения workspace из automation-контура.

## Рекомендуемая рабочая позиция до следующего уточнения

- считать, что Stage 2 заканчивается серверным Telegram webhook capture-процессом;
- считать, что durable Obsidian workspace живет в отдельном git-репозитории;
- не считать, что Codex обязан постоянно крутиться на сервере для Stage 2;
- не считать Stage 2 зависящим от Obsidian CLI и skills;
- вернуться к этим вопросам при планировании runtime implementation и Stage 3-4.
