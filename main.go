package main

import (
	"net/http"
)

func main() {
	Printfln("Starting HTTP Server")
	http.ListenAndServe(":5000", nil)
}
