package main

import (
	"html/template"
	"os"
)

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}
func main() {
	
	allTemplates, err := template.ParseFiles("templates/template.html","templates/list.html")



	if err == nil {
		selectedTemplated := allTemplates.Lookup("mainTemplate")
		err = Exec(selectedTemplated)
	}
	if err != nil {
		Printfln("Error: %v %v", err.Error())
	}
}
