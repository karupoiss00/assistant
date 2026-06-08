package workspace

import "context"

type MessageKind string

const (
	MessageKindText  MessageKind = "text"
	MessageKindVoice MessageKind = "voice"
)

type Entry struct {
	EntryID           string
	TelegramUpdateID  string
	TelegramMessageID string
	MessageKind       MessageKind
	StoredText        string
	TranscriptText    string
	ProcessingStatus  string
	ReceivedAt        string
}

type Writer struct {
	workspaceRootPath     string
	dailyInboxPathPattern string
}

func NewWriter(workspaceRootPath, dailyInboxPathPattern string) *Writer {
	return &Writer{
		workspaceRootPath:     workspaceRootPath,
		dailyInboxPathPattern: dailyInboxPathPattern,
	}
}

func (w *Writer) AppendInboxEntry(ctx context.Context, entry Entry) error {
	_ = ctx
	_ = entry
	return nil
}

func (w *Writer) UpdateVoiceTranscript(ctx context.Context, entryID, transcript, status string) error {
	_ = ctx
	_ = entryID
	_ = transcript
	_ = status
	return nil
}
