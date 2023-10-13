package main
import (
	"net/http"
	"os"
	"time"
	"io"
)
func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)
	response, err := http.Get("http://localhost:5000/html")
	if (err == nil && response.StatusCode == http.StatusOK) {
		data, err := io.ReadAll(response.Body)
		if (err == nil) {
			defer response.Body.Close()
			os.Stdout.Write(data)
		}
	} else {
		Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
	}
}