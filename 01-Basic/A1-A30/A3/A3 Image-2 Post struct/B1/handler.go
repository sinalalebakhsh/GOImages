package cms

import (
	"net/http"
	"strings"
	"time"
)

// ServePage serves a page based on the route mathed.
// This will match any URL beginning with /page
func ServePage(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/page/")

	if path == "" {
		http.NotFound(w, r)
		return
	}

	p := &Page{
		Title:   strings.ToTitle(path),
		Content: "Here is my page",
	}

	Templ.ExecuteTemplate(w, "page", p)

}

// ServePost servers a post
func ServePost(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/post/")

	if path == "" {
		http.NotFound(w, r)
		return
	}

	p := &Post{
		Title:   strings.ToTitle(path),
		Content: "Here is my page",
		Comments: []*Comment{
			&Comment{
				Author:        "Sina Sina LALE LALE",
				Comment:       "This is Great!",
				DatePublished: time.Now(),
			},
		},
	}

	Templ.ExecuteTemplate(w, "post", p)
}

// HandleNew handles preview logic
func HandlesNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Templ.ExecuteTemplate(w, "new", nil)
	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()

		if contentType == "page" {
			Templ.ExecuteTemplate(w, "page", &Page{
				Title:   title,
				Content: content,
			})
			return
		}
		if contentType == "post" {
			Templ.ExecuteTemplate(w, "post", &Post{
				Title:   title,
				Content: content,
			})
			return
		}
	default:
		http.Error(w, "Method not supported: "+r.Method, http.StatusMethodNotAllowed)
	}
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Go Projects CMS",
		Content: "Welcome to our home page",
		Posts: []*Post{
			&Post{
				Title:         "Hello, World",
				Content:       "Thanks for learning with me",
				DatePublished: time.Now(),
			},
			&Post{
				Title:         "Another Post title",
				Content:       "For Cheking attract comments",
				DatePublished: time.Now().Add(-time.Hour),
				Comments: []*Comment{
					&Comment{
						Author:        "Sina Laleh Bakhsh",
						Comment:       "Nevermind, its just commented on Test Post...",
						DatePublished: time.Now().Add(-time.Hour / 2),
					},
				},
			},
		},
	}

	Templ.ExecuteTemplate(w, "page", p)
}
