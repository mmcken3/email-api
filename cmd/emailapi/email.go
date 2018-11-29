package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/mmcken3/email-api/internal/contact"
	"github.com/mmcken3/email-api/internal/gmail"
)

func emailHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling email request")

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

	sendEmail := gmail.Email{
		UserName:    cfg.UserName,
		Password:    cfg.Password,
		Server:      cfg.Server,
		Port:        cfg.Port,
		SendTo:      []string{cfg.SendTo},
		FromAddress: cfg.UserName,
	}

	err = sendEmail.SendEmail(m.Name, m.Email, m.Message)
	if err != nil {
		log.Println("err : ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		render.JSON(w, r, resp{Message: "failure"})
		return
	}

	log.Println("Email sent success")
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, resp{Message: "success"})
}
