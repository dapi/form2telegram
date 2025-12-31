package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dapi/form2telegram/internal/formatter"
)

type MessageSender interface {
	SendMessage(text string) error
}

type WebhookHandler struct {
	sender MessageSender
}

func NewWebhookHandler(sender MessageSender) *WebhookHandler {
	return &WebhookHandler{sender: sender}
}

func (h *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var form formatter.FormData
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Printf("Failed to decode request: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	message := formatter.FormatForm(&form)
	if err := h.sender.SendMessage(message); err != nil {
		log.Printf("Failed to send message: %v", err)
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
