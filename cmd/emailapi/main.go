package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	UserName        string `envconfig:"EMAIL_USER"`
	Password        string `envconfig:"EMAIL_PASSWORD"`
	Server          string `envconfig:"EMAIL_SERVER"`
	Port            string `envconfig:"EMAIL_PORT"`
	SendToM         string `envconfig:"EMAIL_M"`
	SendToK         string `envconfig:"EMAIL_K"`
	FromAddress     string `envconfig:"EMAIL_FROM"`
	TwilioSID       string `envconfig:"TWILIO_ACCOUNT_SID" required:"true"`
	TwilioAuthToken string `envconfig:"TWILIO_AUTH_TOKEN" required:"true"`
}

var cfg config

func main() {
	log.Println("Starting email api!")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := envconfig.Process("import", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health Check OK"))
	})
	// r.Post("/v1/send/email", emailHandler)
	r.Post("/v1/send/text", textHandler)

	http.ListenAndServe(":3000", r)
}
