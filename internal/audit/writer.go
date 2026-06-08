package audit

import "context"

type Event struct {
	AuditID          string
	TelegramUpdateID string
	EntryID          string
	AttemptID        string
	EventType        string
	Summary          string
}

type Writer struct {
	workspaceRootPath string
	logPathPattern    string
}

func NewWriter(workspaceRootPath, logPathPattern string) *Writer {
	return &Writer{
		workspaceRootPath: workspaceRootPath,
		logPathPattern:    logPathPattern,
	}
}

func (w *Writer) AppendEvent(ctx context.Context, event Event) error {
	_ = ctx
	_ = event
	return nil
}
