package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/mmcken3/email-api/internal/gmail"
	"github.com/mmcken3/email-api/internal/twilio"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Config is the config for the handler.
type Config struct {
	Logger       *logrus.Logger
	TwilioConfig twilio.Config
	EmailHandler gmail.Email
}

// Handler is the global handler for the api.
type Handler struct {
	http.Handler

	emailHandler gmail.Email
	twilioConfig twilio.Config
	logger       *logrus.Logger
}

func isValidConfig(c Config) error {
	if c.Logger == nil {
		return errors.New("logger cannot be nil")
	}
	return nil
}

// New returns a new handler.
func New(c Config) (*Handler, error) {
	if err := isValidConfig(c); err != nil {
		return nil, errors.Wrap(err, "invalid handler config")
	}

	h := Handler{
		logger:       c.Logger,
		twilioConfig: c.TwilioConfig,
		emailHandler: c.EmailHandler,
	}

	r := chi.NewRouter()

	// Middleware set up
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Recoverer)

	r.Route("/", func(r chi.Router) {
		// set up routes
		r.Post("/v1/send/email", h.sendEmailHandler)
		r.Post("/v1/send/text", h.sendTextHandler)
	})
	r.Get("/health", h.health)

	h.Handler = r
	return &h, nil
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	// Add any DB, Redis, or server pings here to have a full health check.
	render.JSON(w, r, struct {
		Health string `json:"health"`
	}{
		Health: "OK",
	})
}
