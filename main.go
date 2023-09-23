package main

import (
	"html/template"
	"os"
)

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}
func main() {
	allTemplates, err := template.ParseGlob("templates/*.html")
	if err == nil {
		selectedTemplated := allTemplates.Lookup("template.html")
		err = Exec(selectedTemplated)
	}

	
	if err != nil {
		Printfln("Error: %v %v", err.Error())
	}
}
