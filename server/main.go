package main

import (
	"log"
	"net/http"

	"github.com/itsHabib/twiliogo"
)

func main() {
	http.Handle("/api/google-home-sms/", &twiliogo.TwilioHandler{})
	log.Print("Server starting on port 8081\n")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
