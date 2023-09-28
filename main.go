package main
import (
	"net/http"
	"io"
)
type StringHandler struct {
	message string
}
func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
		request *http.Request) {
	io.WriteString(writer, sh.message)
}
func main() {
	err := http.ListenAndServe(":5000", StringHandler{ message: "Hello, World"})
	if (err != nil) {
		Printfln("Error: %v", err.Error())
	}
}