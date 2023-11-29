package main

import (
	"net/http"
	"web3/pkg/config"
	"web3/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

/* install go-chi
https://github.com/go-chi/chi#install

	go get -u github.com/go-chi/chi/v5

*/

func routes(app *config.AppConfig) http.Handler {
	// Mux stand for HTTP Request Multiplexer
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(LosgRequestInfo)
	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	return mux
}




