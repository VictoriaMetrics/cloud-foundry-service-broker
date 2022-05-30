package httprouter

import (
	"fmt"
	"github.com/VictoriaMetrics/cloud-foundry-service-broker/lib/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
)

// NewRouter creates chi router instance
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(logger.ChiLogger(logger.DefaultLogger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	return r
}

// GetBindAddr get bind address for http server
func GetBindAddr() string {
	return fmt.Sprintf("%s:%d", *bindHost, *bindPort)
}
