package main

import (
	"html/template"
)

// Templ is an exported variable. 
// Templ is stand for Templates
var Templ = template.Must(template.ParseGlob("/templates/*"))

func main()  {
}