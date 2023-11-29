package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"web3/pkg/config"
	"web3/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// Mux stand for HTTP Request Multiplexer
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(LosgRequestInfo)

	mux.Use(NoSurf)
	mux.Use(SetupSession)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	return mux
}
