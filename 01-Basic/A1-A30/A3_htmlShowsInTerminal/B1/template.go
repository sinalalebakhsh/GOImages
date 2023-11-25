package cms

import (
	"html/template"
)

// Templ is an exported variable. 
// Templ is stand for Templates
var Templ = template.Must(template.ParseGlob("../templates/*"))


type Page struct {
	Title string
	Content string
}