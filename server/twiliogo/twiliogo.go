package twiliogo

import (
	"net/http"
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

}
