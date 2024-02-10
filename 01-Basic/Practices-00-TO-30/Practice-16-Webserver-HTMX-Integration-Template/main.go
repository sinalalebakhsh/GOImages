package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main()  {
	fmt.Println("Sina Lalehbakhsh")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Sina Homepage\n")
		io.WriteString(w, "Request Method: ")
		io.WriteString(w, r.Method)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8765", nil))

}