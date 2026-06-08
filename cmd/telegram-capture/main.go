package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"assistant/internal/audit"
	"assistant/internal/capture"
	"assistant/internal/config"
	"assistant/internal/telegram"
	"assistant/internal/transcribe"
	"assistant/internal/workspace"
)

func main() {
	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	workspaceWriter := workspace.NewWriter(cfg.WorkspaceRootPath, cfg.DailyInboxPathPattern)
	auditWriter := audit.NewWriter(cfg.WorkspaceRootPath, cfg.CaptureAuditLogPathPattern)
	transcriber := transcribe.NewNoopClient()
	service := capture.NewService(cfg, workspaceWriter, auditWriter, transcriber)
	handler := telegram.NewWebhookHandler(cfg.TelegramWebhookPath, service)

	server := &http.Server{
		Addr:              cfg.ListenAddr,
		Handler:           handler.Routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("telegram capture listening on %s", cfg.ListenAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown: %v", err)
	}
}
