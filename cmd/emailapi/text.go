package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/mmcken3/email-api/internal/twilio"
)

type message struct {
	Name    string `json:"name"`
	Email   string `json:"email_address"`
	Message string `json:"message"`
}

func textHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling email request")

	var m message
	err := json.NewDecoder(r.Body).Decode(&m)
	r.Body.Close()
	if err != nil {
		log.Println("err : ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	site := "mitchell"
	if strings.Contains(r.Referer(), "katie") {
		site = cfg.SendToK
	}

	twilio.SendTextMessage(m.Name, m.Email, m.Message, site, cfg.TwilioSID, cfg.TwilioAuthToken)
	log.Println("Message sent")

	w.WriteHeader(http.StatusOK)
}
