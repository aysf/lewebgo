package main

import (
	"net/http"

	"github.com/aysf/lewebgo/config"
	"github.com/aysf/lewebgo/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Route(app *config.AppConfig) http.Handler {

	mux := chi.NewMux()

	mux.Use(MyCustomlogger)
	mux.Use(middleware.Logger)

	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
