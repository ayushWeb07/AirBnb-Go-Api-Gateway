package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RouterInterface interface {
	Register(r *chi.Mux)
}

func SetupRouter() *chi.Mux {
	// create the router instance
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// setup health routes
	SetupHealthRouter(router)

	return router
}
