package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/mmcken3/email-api/cmd/emailapi/handler"
	"github.com/mmcken3/email-api/internal/gmail"
	"github.com/mmcken3/email-api/internal/twilio"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type config struct {
	UserName string `envconfig:"EMAIL_USER"`
	Password string `envconfig:"EMAIL_PASSWORD"`
	Server   string `envconfig:"EMAIL_SERVER"`
	Port     string `envconfig:"EMAIL_PORT"`
	SendTo   string `envconfig:"SEND_TO"`

	TwilioSID       string `envconfig:"TWILIO_ACCOUNT_SID" required:"true"`
	TwilioAuthToken string `envconfig:"TWILIO_AUTH_TOKEN" required:"true"`
	FromNumber      string `envconfig:"FROM_TWILIO_NUMBER" required:"true"`
	ToNumber        string `envconfig:"TO_PHONE_NUMBER" required:"true"`

	Debug bool `envconfig:"DEBUG" default:"true"`
}

var cfg config
var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Println("Starting email api!")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := envconfig.Process("import", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	if !cfg.Debug {
		log.SetFormatter(&logrus.JSONFormatter{})
	}
}

func main() {
	log.Println("Email API starting up!")

	cert, err := loadCert()
	if err != nil {
		log.Fatal(errors.Wrap(err, "loading cert"))
	}

	// set up our global handler
	h, err := handler.New(handler.Config{
		Logger: log,
		TwilioConfig: twilio.Config{
			SID:        cfg.TwilioSID,
			Token:      cfg.TwilioAuthToken,
			ToNumber:   cfg.ToNumber,
			FromNumber: cfg.FromNumber,
		},
		EmailHandler: gmail.Email{
			UserName:    cfg.UserName,
			Password:    cfg.Password,
			Server:      cfg.Server,
			Port:        cfg.Port,
			SendTo:      []string{cfg.SendTo},
			FromAddress: cfg.UserName,
		},
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "handler new"))
	}

	server := &http.Server{
		Handler: h,
		Addr:    fmt.Sprintf(":%d", 3000),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	// do graceful server shutdown
	go gracefulShutdown(server, time.Second*30)

	log.Infof("listening on port %d", 3000)
	if err := server.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
		log.WithError(err).Fatal("cannot start a server")
	}
}

// gracefulShutdown shuts down our server in a graceful way.
func gracefulShutdown(server *http.Server, timeout time.Duration) {
	sigStop := make(chan os.Signal)

	signal.Notify(sigStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	<-sigStop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.WithError(err).Fatal("graceful shutdown failed")
	}

	log.Info("graceful shutdown complete")
}
