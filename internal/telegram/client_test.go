package telegram

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_SendMessage_Success(t *testing.T) {
	var receivedRequest struct {
		ChatID    string `json:"chat_id"`
		Text      string `json:"text"`
		ParseMode string `json:"parse_mode"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/botTEST_TOKEN/sendMessage" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		if err := json.NewDecoder(r.Body).Decode(&receivedRequest); err != nil {
			t.Fatalf("failed to decode request: %v", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ok": true}`))
	}))
	defer server.Close()

	client := NewClient("TEST_TOKEN", "12345")
	client.baseURL = server.URL

	err := client.SendMessage("Hello, World!")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if receivedRequest.ChatID != "12345" {
		t.Errorf("chat_id = %q, want %q", receivedRequest.ChatID, "12345")
	}
	if receivedRequest.Text != "Hello, World!" {
		t.Errorf("text = %q, want %q", receivedRequest.Text, "Hello, World!")
	}
	if receivedRequest.ParseMode != "Markdown" {
		t.Errorf("parse_mode = %q, want %q", receivedRequest.ParseMode, "Markdown")
	}
}

func TestClient_SendMessage_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"ok": false, "description": "Bad Request: chat not found"}`))
	}))
	defer server.Close()

	client := NewClient("TEST_TOKEN", "12345")
	client.baseURL = server.URL

	err := client.SendMessage("Hello")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestClient_SendMessage_NetworkError(t *testing.T) {
	client := NewClient("TEST_TOKEN", "12345")
	client.baseURL = "http://localhost:1" // port that doesn't exist

	err := client.SendMessage("Hello")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
