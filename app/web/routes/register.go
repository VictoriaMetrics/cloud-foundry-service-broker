package routes

import (
	"github.com/VictoriaMetrics/cloud-foundry-service-broker/app/web/controllers"
	"github.com/go-chi/chi/v5"
)

// Register registers all the routes for application
func Register(r *chi.Mux) {
	r.Get("/", controllers.HelloHandler)
}
