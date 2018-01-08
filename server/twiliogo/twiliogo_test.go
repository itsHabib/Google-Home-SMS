package twiliogo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestTwilioHandlerWithMissingParameters ensures that the handler is
// called with the correct parameters
func TestTwilioHandlerWithMissingParameters(t *testing.T) {
	tests := []twilioSmsRequest{
		{From: "18888888888"},
		{To: "18888888888"},
		{Body: "Missing params"},
		{From: "18888888888", Body: "I'm missing a to param"},
		{To: "18888888888", Body: "I'm missing a from param"},
		{To: "18888888888", From: "18888888888"},
	}
	handler := TwilioHandler{}

	for _, tt := range tests {
		data, err := json.Marshal(tt)
		if err != nil {
			t.Fatal(err)
		}
		reader := bytes.NewReader(data)
		req := httptest.NewRequest("POST", "/api/google-home-sms", reader)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Bad request, missing parameters, "+
				"expected code=%v, got=%v", http.StatusBadRequest, rr.Code)
		}
	}
}

// TestTwilioHandlerWithWrongMethods ensures that the handler is
// called with the correct method, POST
func TestTwilioHandlerWithWrongMethods(t *testing.T) {
	requestData := twilioSmsRequest{
		From: "18888888888",
		To:   "18888888888",
		Body: "Wrong method testing request",
	}
	testMethods := []string{"GET", "OPTIONS", "HEAD", "DELETE", "CONNECT", "PUT"}
	data, err := json.Marshal(requestData)
	if err != nil {
		t.Fatal(err)
	}

	handler := TwilioHandler{}
	reader := bytes.NewReader(data)

	for _, tt := range testMethods {
		req := httptest.NewRequest(tt, "/api/google-home-sms/", reader)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("Method not allowed, expected code=%v, got=%v",
				http.StatusMethodNotAllowed, rr.Code)
		}
	}
}

func TestSendSMSWithNoAuth(t *testing.T) {
	tests := []TwilioCilent{
		{AccountSID: ""},
		{AuthToken: ""},
	}
	smsRequest := twilioSmsRequest{
		From: "18888888888",
		To:   "18888888888",
		Body: "body",
	}
	for _, tt := range tests {
		err := tt.SendSMS(smsRequest)
		if err == nil {
			t.Errorf("No auth provided, expected err=%v, got=%v", "No auth provided", err)
		}
	}
}
