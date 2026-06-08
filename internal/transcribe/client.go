package transcribe

import (
	"context"
	"errors"
)

var ErrNotImplemented = errors.New("transcription client not implemented")

type Client interface {
	Transcribe(ctx context.Context, filePath string) (string, error)
}

type NoopClient struct{}

func NewNoopClient() NoopClient {
	return NoopClient{}
}

func (NoopClient) Transcribe(ctx context.Context, filePath string) (string, error) {
	_ = ctx
	_ = filePath
	return "", ErrNotImplemented
}
