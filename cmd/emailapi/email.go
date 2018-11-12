package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	"github.com/mmcken3/email-api/internal/contact"
	"github.com/pkg/errors"
)

// Email is a struct used for sending go emails.
type Email struct {
	UserName    string
	Password    string
	Server      string
	Port        string
	SendTo      []string
	FromAddress string
}

func emailHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling email request")

	toSend := cfg.SendToM
	if strings.Contains(r.Referer(), "katie") {
		toSend = cfg.SendToK
	}

	sendEmail := Email{
		UserName:    cfg.UserName,
		Password:    cfg.Password,
		Server:      cfg.Server,
		Port:        cfg.Port,
		SendTo:      []string{toSend},
		FromAddress: cfg.FromAddress,
	}

	log.Println(sendEmail)

	w.WriteHeader(http.StatusOK)
}

// SendEmail sends an email to email address e.
func (se *Email) SendEmail(contact contact.Contact) error {
	var b bytes.Buffer

	b.Write([]byte("To: "))

	for _, email := range se.SendTo {
		b.Write([]byte(email + ", "))
	}

	b.Write([]byte("\r\nSubject: Contact From Website"))
	b.Write([]byte(time.Now().Format("Jan-01-06 03:04 PM") + "\r\n"))
	b.Write([]byte("\r\nType: "))
	b.Write([]byte(("Stuff is here") + "\n"))
	b.Write([]byte("\n\n"))
	fmt.Println(b.String())

	//msg := []byte("\r\nSubject: CU Fix It Request\r\nMessage Content Here")

	// Set up authentication information
	auth := smtp.PlainAuth("", se.UserName, se.Password, se.Server)

	msg := b.Bytes()
	err := smtp.SendMail(se.Server+":"+se.Port, auth, se.FromAddress, se.SendTo, msg)
	if err != nil {
		return errors.Wrapf(err, "Failed when sending email.")
	}
	return nil
}
