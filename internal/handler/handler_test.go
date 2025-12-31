package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockSender struct {
	lastMessage string
	err         error
}

func (m *mockSender) SendMessage(text string) error {
	m.lastMessage = text
	return m.err
}

func TestWebhookHandler_Success(t *testing.T) {
	sender := &mockSender{}
	h := NewWebhookHandler(sender)

	body := `{"answers":[{"key":"email","value":"test@example.com"}]}`
	req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if sender.lastMessage != "*email:* test@example.com" {
		t.Errorf("message = %q, want %q", sender.lastMessage, "*email:* test@example.com")
	}
}

func TestWebhookHandler_InvalidJSON(t *testing.T) {
	sender := &mockSender{}
	h := NewWebhookHandler(sender)

	req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBufferString("not json"))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestWebhookHandler_WrongMethod(t *testing.T) {
	sender := &mockSender{}
	h := NewWebhookHandler(sender)

	req := httptest.NewRequest(http.MethodGet, "/webhook", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
}

func TestWebhookHandler_TelegramError(t *testing.T) {
	sender := &mockSender{err: http.ErrHandlerTimeout}
	h := NewWebhookHandler(sender)

	body := `{"answers":[{"key":"test","value":"value"}]}`
	req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusInternalServerError)
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	HealthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if rec.Body.String() != "OK" {
		t.Errorf("body = %q, want %q", rec.Body.String(), "OK")
	}
}
