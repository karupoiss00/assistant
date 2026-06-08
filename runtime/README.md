# Runtime

В этой папке можно хранить только временные или служебные файлы рантайма.

По умолчанию содержимое этой папки не считается источником истины. Источником
истины остаются production workspace, конфигурация и репозиторные audit
records.

## Stage 2 Temporary Voice Artifact Lifecycle

- Временные voice artifacts разрешены только как краткоживущие файлы внутри
  `runtime/`.
- Runtime может скачать voice-файл во временный путь вида
  `runtime/<run_id>/<telegram_file_id>.ogg`.
- Такой файл существует только на время транскрибации и cleanup-handling.
- После success path или failure path временный файл должен быть удален.
- Temporary voice artifact не может считаться durable state и не должен
  становиться частью production workspace по умолчанию.
- Retry должен использовать durable inbox context и audit trail, а не требовать
  сохранения старого временного voice-файла.
