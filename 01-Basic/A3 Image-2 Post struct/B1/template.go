package cms

import (
	"html/template"
	"time"
)

// Templ is an exported variable. 
// Templ is stand for Templates
var Templ = template.Must(template.ParseGlob("../templates/*"))


type Page struct {
	Title string
	Content string
	Posts []*Post
}

type Post struct {
	Title string
	Content string
	DatePublished time.Time
	Comments []*Comment
}

type Comment struct {
	Author string
	Comment string
	DatePublished time.Time	
}