package twiliogo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSMSHandlerWithMissingParameters(t *testing.T) {
	tests := []twilioSmsRequest{
		{From: "18186247532"},
		{To: "18186247532"},
		{Body: "Missing params"},
		{From: "18186247532", Body: "I'm missing a to param"},
		{To: "18186247532", Body: "I'm missing a from param"},
		{To: "18186247532", From: "18186247532"},
	}
	handler := TwilioHandler{}

	for tt := range tests {
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
