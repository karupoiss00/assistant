package capture

import (
	"context"

	"assistant/internal/audit"
	"assistant/internal/config"
	"assistant/internal/transcribe"
	"assistant/internal/workspace"
)

type Service struct {
	config     config.Config
	workspace  *workspace.Writer
	audit      *audit.Writer
	transcribe transcribe.Client
}

func NewService(cfg config.Config, workspaceWriter *workspace.Writer, auditWriter *audit.Writer, transcriber transcribe.Client) *Service {
	return &Service{
		config:     cfg,
		workspace:  workspaceWriter,
		audit:      auditWriter,
		transcribe: transcriber,
	}
}

type TextInput struct {
	TelegramUpdateID  string
	TelegramMessageID string
	ChatID            string
	SenderID          string
	Text              string
	ReceivedAt        string
}

type VoiceInput struct {
	TelegramUpdateID  string
	TelegramMessageID string
	ChatID            string
	SenderID          string
	FileID            string
	ReceivedAt        string
}

func (s *Service) HandleText(ctx context.Context, input TextInput) error {
	_ = ctx
	_ = input
	return nil
}

func (s *Service) HandleVoice(ctx context.Context, input VoiceInput) error {
	_ = ctx
	_ = input
	return nil
}
