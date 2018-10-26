package homepage

import (
	"log"
	"net/http"
)

const message = "Hello, world!"

type Handlers struct {
	logger * log.Logger
}

func (h * Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/home", h.Home)
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func New(logger *log.Logger) *Handlers {
	return &Handlers{logger: logger}
}