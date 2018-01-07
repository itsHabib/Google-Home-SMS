package twiliogo

import (
	"encoding/json"
	"net/http"
)

const (
	TWILIO_URL = "https://api.twilio.com/2010-04-01/Accounts"
)

type TwilioCilent struct {
	AccountSID string
	AuthToken  string
}

type TwilioHandler struct{}

type twilioSmsRequest struct {
	To   string
	From string
	Body string
}

func (tw *TwilioHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	requestData := new(twilioSmsRequest)
	err := decoder.Decode(requestData)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	if requestData.Body == "" || requestData.From == "" || requestData.To == "" {
		http.Error(w, "Bad Request, missing parameters", http.StatusBadRequest)
	}
}
