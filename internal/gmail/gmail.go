package gmail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"time"

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

// SendEmail sends an email to email address e.
// WIP will probably change as gmail has improved email stuff now from
// go code.
func (se *Email) SendEmail(name string, email string, message string) error {
	var b bytes.Buffer

	b.Write([]byte("To: "))

	for _, email := range se.SendTo {
		b.Write([]byte(email + ", "))
	}

	nameString := fmt.Sprintf("You have been contacted by %v\n", name)
	emailString := fmt.Sprintf("Contact them back at %v\n", email)
	messageString := fmt.Sprintf("The message is: %v\n", message)

	b.Write([]byte("\r\nSubject: Contact From Personal Website"))
	b.Write([]byte(time.Now().Format("Jan-01-06 03:04 PM") + "\r\n"))
	b.Write([]byte(("\r\n")))
	b.Write([]byte((nameString)))
	b.Write([]byte((emailString)))
	b.Write([]byte((messageString)))
	b.Write([]byte("\n\n"))

	//Set up authentication information
	auth := smtp.PlainAuth("", se.UserName, se.Password, se.Server)

	msg := b.Bytes()
	err := smtp.SendMail(se.Server+":"+se.Port, auth, se.FromAddress, se.SendTo, msg)
	if err != nil {
		return errors.Wrapf(err, "Failed when sending email.")
	}
	return nil
}
