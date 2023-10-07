package main
import (
	"fmt"
	"io"
	"net/http"
)
func HandleMultipartForm(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Name: %v, City: %v\n", request.FormValue("name"),
		request.FormValue("city"))
	fmt.Fprintln(writer, "------")
	file, header, err := request.FormFile("files")
	if err == nil {
		defer file.Close()
		fmt.Fprintf(writer, "Name: %v, Size: %v\n", header.Filename, header.Size)
		for k, v := range header.Header {
			fmt.Fprintf(writer, "Key: %v, Value: %v\n", k, v)
		}
		fmt.Fprintln(writer, "------")
		io.Copy(writer, file)
	} else {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
func init() {
	http.HandleFunc("/forms/upload", HandleMultipartForm)
}
