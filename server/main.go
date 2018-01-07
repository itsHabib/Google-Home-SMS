package main

import (
	"log"
	"net/http"

	"github.com/itsHabib/google-home-sms/server/twiliogo"
)

func main() {
	err := http.ListenAndServe(":8081", nil)
	http.Handle("/api/google-home-sms/", &twiliogo.TwilioHandler{})
	if err != nil {
		log.Fatal(err)
	}
}
