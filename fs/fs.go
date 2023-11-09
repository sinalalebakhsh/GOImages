package main

import (
	"net/http"
	"os"
)

func main()  {
	dir, _ := os.Getwd()
	http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
}