package main

import (
	"log"
	"net/http"
	"web3/pkg/config"
	"web3/pkg/handlers"
)

func main()  {
	var app config.AppConfig

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr: ":8080",
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

/* 
For running in current directory , Open Terminal and write:
	
	go run ./cmd/web
	
*/