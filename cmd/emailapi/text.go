package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/mmcken3/email-api/internal/contact"
	"github.com/mmcken3/email-api/internal/twilio"
)

// resp is a struct to be used as the json reponse holding a message.
type resp struct {
	Message string `json:"message"`
}

// textHandler will send a text to the configured number using the configured
// twilio account or return a failed response
func textHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling text request")

	// Decode the contact from the POST body
	var m contact.Contact
	err := json.NewDecoder(r.Body).Decode(&m)
	r.Body.Close()
	if err != nil {
		log.Println("err : ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		render.JSON(w, r, resp{Message: "failure"})
		return
	}

	site := r.Referer()

	// Create message body
	textMsgBody := fmt.Sprintf("You have been contacted by %v. %v. Message: %v. Site: %v", m.Name, m.Email, m.Message, site)
	twilioConfig := twilio.Config{
		Message:    textMsgBody,
		SID:        cfg.TwilioSID,
		Token:      cfg.TwilioAuthToken,
		ToNumber:   cfg.ToNumber,
		FromNumber: cfg.FromNumber,
	}

	// send message using twilio package
	twilio.SendTextMessage(twilioConfig)

	log.Println("Message sent sucess")
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, resp{Message: "success"})
}
