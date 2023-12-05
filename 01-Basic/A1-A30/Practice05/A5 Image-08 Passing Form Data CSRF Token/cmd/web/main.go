// Teacher: derek banas

package main

import (
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
	"web3/pkg/config"
	"web3/pkg/handlers"
)

var sessionManager *scs.SessionManager
var app config.AppConfig

func main() {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = false
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = sessionManager

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	/* Previous development
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
	*/
}


/* installations 

1-go-chi
https://github.com/go-chi/chi#install
	go get -u github.com/go-chi/chi/v5

2-CSRF Token
	go get github.com/justinas/nosurf
	
*/


/*
For running in current directory , Open Terminal and write:

	go run ./cmd/web

*/
