package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/scottywm/bookings/pkg/config"
	"github.com/scottywm/bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// this feature is in the chi library and what it does is, if an error happens inside our application, instead of dying, it prints the error and keeps the application running.
	mux.Use(middleware.Recoverer)
	// this is how we actually create the cookie for each request with the settings stipulated in the file middleware
	mux.Use(NoSurf)
	// this enables our server to save the session data to memory and send the session data inside a cookie via requests going to and from the client.
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
