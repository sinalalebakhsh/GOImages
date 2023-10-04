package main
import (
	"net/http"
	"io"
)
type StringHandler struct {
	message string
}
func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if (request.URL.Path == "/favicon.ico") {
		Printfln("Request for icon detected - returning 404")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
}
func main() {
	err := http.ListenAndServe(":5000", StringHandler{ message: "Hello, World"})
	if (err != nil) {
		Printfln("Error: %v", err.Error())
	}
}
