package twiliogo

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	TWILIO_URL = "https://api.twilio.com/2010-04-01/Accounts"
)

// TwilioClient represents a client for the Twilio API
type TwilioClient struct {
	AccountSID string
	AuthToken  string
}

// TwilioHandler is the handler struct responsible for serving http requests
type TwilioHandler struct{}

type twilioSmsRequest struct {
	To   string
	From string
	Body string
}

// Handles the data sent in a POST request to /api/google-home-sms/
func (twh *TwilioHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, "{sent: false, time: 0, bad_request: Method not allowed}",
			http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "{sent: false, time: 0, bad_request: Content type must be application/json}",
			http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	requestData := new(twilioSmsRequest)
	err := decoder.Decode(requestData)
	if err != nil {
		http.Error(w, fmt.Sprintf("{sent: false, time: 0, bad_request: bad request, err: %s}", err.Error()), http.StatusBadRequest)
		return
	}
	if requestData.Body == "" || requestData.From == "" || requestData.To == "" {
		http.Error(w, fmt.Sprintf("{sent: false, time: 0, bad_request: missing parameters, err: %s}", err.Error()), http.StatusBadRequest)
		return
	}
	twClient := &TwilioClient{
		AccountSID: os.Getenv("TWILIO_SID"),
		AuthToken:  os.Getenv("TWILIO_TOKEN"),
	}
	err = twClient.SendSMS(*requestData)
	if err != nil {
		http.Error(w, fmt.Sprintf("{sent:false, time_stamp:0, err: %s}", err.Error()), http.StatusInternalServerError)
		return
	}
	jsonResponse := struct {
		Sent      bool      `json:"sent"`
		TimeStamp time.Time `json:"time_stamp"`
	}{
		Sent:      true,
		TimeStamp: time.Now(),
	}
	json.NewEncoder(w).Encode(jsonResponse)

}

// SendSMS is responsible for sending the POST request to twilio to actually
// send the sms
func (twc *TwilioClient) SendSMS(smsData twilioSmsRequest) error {
	if twc.AuthToken == "" || twc.AccountSID == "" {
		return errors.New("No auth provided")
	}

	accountSID := os.Getenv("TWILIO_SID")
	authToken := os.Getenv("TWILIO_TOKEN")
	urlStr := TWILIO_URL + "/" + accountSID + "/Messages.json"

	data := url.Values{}
	data.Set("To", smsData.To)
	data.Set("From", smsData.From)
	data.Set("Body", smsData.Body)
	dataReader := *strings.NewReader(data.Encode())

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlStr, &dataReader)
	if err != nil {
		return err
	}
	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		err = nil
	}
	return err
}
