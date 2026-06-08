package config

import (
	"errors"
	"os"
)

const (
	defaultListenAddr            = ":8080"
	defaultWebhookPath           = "/telegram/webhook"
	defaultAssistantConfigPath   = "System/assistant-config.yaml"
	defaultDailyInboxPathPattern = "System/Inbox/YYYY/MM/YYYY-MM-DD.md"
	defaultCaptureLogPathPattern = "System/Logs/telegram-capture/YYYY/MM/YYYY-MM-DD.ndjson"
	defaultRuntimeRoot           = "runtime"
)

type Config struct {
	ListenAddr                           string
	TelegramWebhookPath                  string
	WorkspaceRootPath                    string
	AssistantConfigPath                  string
	DailyInboxPathPattern                string
	CaptureAuditLogPathPattern           string
	RuntimeRoot                          string
	VoiceTranscriptionRequiredForConfirm bool
}

func LoadFromEnv() (Config, error) {
	cfg := Config{
		ListenAddr:                           getenv("ASSISTANT_LISTEN_ADDR", defaultListenAddr),
		TelegramWebhookPath:                  getenv("TELEGRAM_WEBHOOK_PATH", defaultWebhookPath),
		WorkspaceRootPath:                    os.Getenv("WORKSPACE_ROOT_PATH"),
		AssistantConfigPath:                  getenv("ASSISTANT_CONFIG_PATH", defaultAssistantConfigPath),
		DailyInboxPathPattern:                getenv("DAILY_INBOX_PATH_PATTERN", defaultDailyInboxPathPattern),
		CaptureAuditLogPathPattern:           getenv("CAPTURE_AUDIT_LOG_PATH_PATTERN", defaultCaptureLogPathPattern),
		RuntimeRoot:                          getenv("RUNTIME_ROOT", defaultRuntimeRoot),
		VoiceTranscriptionRequiredForConfirm: true,
	}

	if cfg.WorkspaceRootPath == "" {
		return Config{}, errors.New("WORKSPACE_ROOT_PATH is required")
	}

	return cfg, nil
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
