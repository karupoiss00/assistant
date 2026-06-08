package telegram

import (
	"encoding/json"
	"net/http"

	"assistant/internal/capture"
)

type WebhookHandler struct {
	webhookPath string
	service     *capture.Service
}

func NewWebhookHandler(webhookPath string, service *capture.Service) *WebhookHandler {
	return &WebhookHandler{
		webhookPath: webhookPath,
		service:     service,
	}
}

func (h *WebhookHandler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", h.handleHealth)
	mux.HandleFunc(h.webhookPath, h.handleWebhook)
	return mux
}

func (h *WebhookHandler) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (h *WebhookHandler) handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload map[string]any
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	_ = payload
	_ = h.service

	w.WriteHeader(http.StatusNoContent)
}
