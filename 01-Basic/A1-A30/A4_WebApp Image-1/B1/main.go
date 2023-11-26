package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter,
		r *http.Request) {
		nb, err := fmt.Fprintf(w, "hello browser")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Bytes written: %d", nb)
	})
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
