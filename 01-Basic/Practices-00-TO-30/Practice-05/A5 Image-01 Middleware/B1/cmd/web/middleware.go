package main

import (
	"fmt"
	"net/http"
	"time"
)

func LosgRequestInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Printf("%d/%d/%d : %d:%d ", 
		now.Month(),
		now.Day(),
		now.Year(),
		now.Hour(),
		now.Minute(),)
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}