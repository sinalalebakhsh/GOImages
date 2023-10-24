package main
import (
	"io"
	"net/http"
	"os"
	"time"
	//"encoding/json"
	//"strings"
	//"net/url"
	"fmt"
	"net/http/cookiejar"
)
func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)
	jar, err := cookiejar.New(nil)
	clients := make([]http.Client, 3)
	for index, client := range clients {
		//jar, err := cookiejar.New(nil)
		if err == nil {
			client.Jar = jar
		}
		for i := 0; i < 3; i++ {
			req, err := http.NewRequest(http.MethodGet,
				"http://localhost:5000/cookie", nil)
			if err == nil {
				response, err := client.Do(req)
				if err == nil && response.StatusCode == http.StatusOK {
					fmt.Fprintf(os.Stdout, "Client %v: ", index)
					io.Copy(os.Stdout, response.Body)
					defer response.Body.Close()
				} else {
					Printfln("Request Error: %v", err.Error())
				}
			} else {
				Printfln("Request Init Error: %v", err.Error())
			}
		}
	}
}