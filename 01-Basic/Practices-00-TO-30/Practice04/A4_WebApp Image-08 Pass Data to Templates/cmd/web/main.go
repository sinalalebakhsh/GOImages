package main

import (
	"log"
	"net/http"
	handlers "web3/pkg/handlers"
)

func main()  {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

/* 
For running in current directory , Open Terminal and write:
	
	go run ./cmd/web
	
*/