#!/usr/bin/env bash

set -euo pipefail

ASSISTANT_HOME="${HOME}/assistant"
CONTROL_DIR="${ASSISTANT_HOME}/control"

echo "[1/6] Creating assistant directories"
mkdir -p \
  "${ASSISTANT_HOME}/pending" \
  "${ASSISTANT_HOME}/logs" \
  "${ASSISTANT_HOME}/scripts" \
  "${ASSISTANT_HOME}/prompts" \
  "${ASSISTANT_HOME}/skills"

echo "[2/6] Checking Node.js"
if command -v node >/dev/null 2>&1; then
  node --version
else
  cat <<'EOF'
Node.js is not installed.

Install Node 24 before continuing. Example for Ubuntu/Debian:
  curl -fsSL https://deb.nodesource.com/setup_24.x | sudo -E bash -
  sudo apt-get install -y nodejs
EOF
  exit 1
fi

echo "[3/6] Checking npm"
if ! command -v npm >/dev/null 2>&1; then
  echo "npm is required but missing."
  exit 1
fi
npm --version

echo "[4/6] Installing OpenClaw"
npm install -g openclaw@latest

echo "[5/6] Preparing local config templates"
if [ -d "${CONTROL_DIR}/deploy/templates" ]; then
  mkdir -p "${HOME}/.openclaw"

  if [ ! -f "${HOME}/.openclaw/openclaw.json" ] && [ -f "${CONTROL_DIR}/deploy/templates/openclaw.json5.template" ]; then
    cp "${CONTROL_DIR}/deploy/templates/openclaw.json5.template" "${HOME}/.openclaw/openclaw.json"
    echo "Created ${HOME}/.openclaw/openclaw.json from template"
  fi

  if [ ! -f "${ASSISTANT_HOME}/config.json" ] && [ -f "${CONTROL_DIR}/deploy/templates/assistant.config.json.template" ]; then
    cp "${CONTROL_DIR}/deploy/templates/assistant.config.json.template" "${ASSISTANT_HOME}/config.json"
    echo "Created ${ASSISTANT_HOME}/config.json from template"
  fi

  if [ ! -f "${ASSISTANT_HOME}/.env" ] && [ -f "${CONTROL_DIR}/deploy/templates/assistant.env.example" ]; then
    cp "${CONTROL_DIR}/deploy/templates/assistant.env.example" "${ASSISTANT_HOME}/.env"
    chmod 600 "${ASSISTANT_HOME}/.env"
    echo "Created ${ASSISTANT_HOME}/.env from template"
  fi
fi

echo "[6/6] Next commands"
cat <<'EOF'
Run these manually:

  openclaw onboard --install-daemon
  openclaw gateway status
  openclaw doctor

Then edit:

  ~/.openclaw/openclaw.json
  ~/assistant/config.json
  ~/assistant/.env
EOF
