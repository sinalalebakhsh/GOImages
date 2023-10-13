package main
import (
	"net/http"
	"os"
	"time"
	"io"
	"encoding/json"
	"strings"
)
func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)
	var builder strings.Builder
	err := json.NewEncoder(&builder).Encode(Products[0])
	if (err == nil) {
		response, err := http.Post("http://localhost:5000/echo", "application/json",strings.NewReader(builder.String()))
		if (err == nil && response.StatusCode == http.StatusOK) {
			io.Copy(os.Stdout, response.Body)
			defer response.Body.Close()
		} else {
			Printfln("Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}
}