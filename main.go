package main
import (
	"net/http"
	"os"
	"io"
	"time"
	//"encoding/json"
	//"strings"
	"net/url"
	//"net/http/cookiejar"
	//"fmt"
)
func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)
	http.DefaultClient.CheckRedirect = func(req *http.Request,
		previous []*http.Request) error {
		if len(previous) == 3 {
			url, _ := url.Parse("http://localhost:5000/html")
			req.URL = url
		}
		return nil
	}
	req, err := http.NewRequest(http.MethodGet,
		"http://localhost:5000/redirect1", nil)
	if (err == nil) {
		var response *http.Response
		response, err = http.DefaultClient.Do(req)
		if (err == nil) {
			io.Copy(os.Stdout, response.Body)
		} else {
			Printfln("Request Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}
}