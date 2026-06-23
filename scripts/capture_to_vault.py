#!/usr/bin/env python3

import argparse
import json
import os
import subprocess
import sys
from dataclasses import dataclass
from datetime import datetime
from pathlib import Path
from zoneinfo import ZoneInfo


@dataclass
class Config:
    vault_path: Path
    pending_path: Path
    logs_path: Path
    inbox_file: str
    timezone: str


def load_config(config_path: Path) -> Config:
    data = json.loads(config_path.read_text())
    return Config(
        vault_path=Path(data["vaultPath"]),
        pending_path=Path(data["pendingPath"]),
        logs_path=Path(data["logsPath"]),
        inbox_file=data["inboxFile"],
        timezone=data["timezone"],
    )


def now_local(tz_name: str) -> datetime:
    return datetime.now(ZoneInfo(tz_name))


def ensure_parent(path: Path) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)


def ensure_inbox(path: Path) -> None:
    ensure_parent(path)
    if not path.exists():
        path.write_text("# Inbox\n\n## Unprocessed\n\n## Processed\n")


def append_under_unprocessed(inbox_path: Path, entry: str) -> None:
    content = inbox_path.read_text()
    marker = "## Unprocessed"
    if marker not in content:
        content = content.rstrip() + "\n\n## Unprocessed\n"
    lines = content.splitlines()
    out = []
    inserted = False
    for idx, line in enumerate(lines):
        out.append(line)
        if line == "## Unprocessed":
            next_line_is_blank = idx + 1 < len(lines) and lines[idx + 1] == ""
            if not next_line_is_blank:
                out.append("")
            out.append(entry)
            inserted = True
    if not inserted:
        out.extend(["", "## Unprocessed", "", entry])
    inbox_path.write_text("\n".join(out).rstrip() + "\n")


def make_entry(kind: str, text: str, ts: datetime) -> str:
    stamp = ts.strftime("%Y-%m-%d %H:%M")
    if kind == "note":
        return f"- {stamp} - {text} #note"
    if kind == "todo":
        return f"- [ ] {text} #todo created:: {stamp}"
    raise ValueError(f"unsupported type: {kind}")


def run_git(args: list[str], repo: Path) -> subprocess.CompletedProcess:
    return subprocess.run(
        ["git", *args],
        cwd=repo,
        text=True,
        capture_output=True,
    )


def append_log(log_path: Path, title: str, body: str) -> None:
    ensure_parent(log_path)
    with log_path.open("a", encoding="utf-8") as fh:
        fh.write(f"[{datetime.utcnow().isoformat()}Z] {title}\n{body}\n\n")


def write_pending(config: Config, reason: str, source: str, original_message: str, ts: datetime) -> Path:
    config.pending_path.mkdir(parents=True, exist_ok=True)
    path = config.pending_path / f"{ts.strftime('%Y-%m-%d-%H-%M-%S')}.md"
    body = (
        "# Pending message\n\n"
        f"created:: {ts.strftime('%Y-%m-%d %H:%M')}\n"
        f"source:: {source}\n"
        f"reason:: {reason}\n\n"
        "## Original message\n\n"
        f"{original_message}\n"
    )
    path.write_text(body)
    return path


def fail(message: str, exit_code: int = 1) -> None:
    print(message)
    sys.exit(exit_code)


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("--type", required=True, choices=["note", "todo"])
    parser.add_argument("--text", required=True)
    parser.add_argument("--source", default="telegram")
    parser.add_argument("--config", default="/root/assistant/config.json")
    args = parser.parse_args()

    raw_text = args.text.strip()
    if not raw_text:
        fail("Empty text.")

    config = load_config(Path(args.config))
    ts = now_local(config.timezone)
    inbox_path = config.vault_path / config.inbox_file
    config.logs_path.mkdir(parents=True, exist_ok=True)
    ensure_inbox(inbox_path)

    original_message = f"{args.type}: {raw_text}"
    entry = make_entry(args.type, raw_text, ts)

    pull = run_git(["pull", "--rebase"], config.vault_path)
    if pull.returncode != 0:
        pending_path = write_pending(config, "git_pull_failed", args.source, original_message, ts)
        append_log(config.logs_path / "git.log", "git pull --rebase failed", pull.stdout + pull.stderr)
        fail(
            "Не смог записать изменение в Obsidian из-за ошибки git pull. "
            f"Сохранил сообщение в pending: {pending_path}"
        )

    append_under_unprocessed(inbox_path, entry)

    status = run_git(["status", "--short"], config.vault_path)
    if status.returncode != 0:
        pending_path = write_pending(config, "git_status_failed", args.source, original_message, ts)
        append_log(config.logs_path / "git.log", "git status failed", status.stdout + status.stderr)
        fail(f"Не удалось проверить состояние git. Сохранил сообщение в pending: {pending_path}")

    add = run_git(["add", config.inbox_file], config.vault_path)
    if add.returncode != 0:
        pending_path = write_pending(config, "git_add_failed", args.source, original_message, ts)
        append_log(config.logs_path / "git.log", "git add failed", add.stdout + add.stderr)
        fail(f"Не удалось добавить изменения в git. Сохранил сообщение в pending: {pending_path}")

    commit = run_git(["commit", "-m", f"assistant: capture {args.type}"], config.vault_path)
    if commit.returncode != 0:
        pending_path = write_pending(config, "git_commit_failed", args.source, original_message, ts)
        append_log(config.logs_path / "git.log", "git commit failed", commit.stdout + commit.stderr)
        fail(f"Не удалось сделать git commit. Сохранил сообщение в pending: {pending_path}")

    push = run_git(["push"], config.vault_path)
    if push.returncode != 0:
        pending_path = write_pending(config, "git_push_failed", args.source, original_message, ts)
        append_log(config.logs_path / "git.log", "git push failed", push.stdout + push.stderr)
        fail(
            "Не смог отправить изменения в git. "
            f"Сохранил сообщение в pending: {pending_path}"
        )

    if args.type == "note":
        print("Записал в inbox.")
    else:
        print("Добавил задачу в inbox.")


if __name__ == "__main__":
    main()
