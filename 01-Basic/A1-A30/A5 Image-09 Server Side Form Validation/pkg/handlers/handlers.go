package handlers

import (
	"log"
	"net/http"
	"web3/models"
	"web3/pkg/config"
	"web3/pkg/forms"
	"web3/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, request *http.Request) {
	m.App.Session.Put(request.Context(), "userid", "sinalalebakhsh")
	render.RenderTemplate(w, "home.page.tmpl", &models.PageData{}, request)
}

func (m *Repository) AboutHandler(w http.ResponseWriter, request *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, "about.page.tmpl", &models.PageData{StrMap: strMap}, request)
	// strMap["title"] = "About Sina Laleh.Bakhsh"
	// strMap["intro"] = "This page is where we talk about ourselves. We love talking about ourselves."
	// userid := m.App.Session.GetString(request.Context(), "userid")
	// strMap["userid"] = userid
}

func (m *Repository) LoginHandler(w http.ResponseWriter, request *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, "login.page.tmpl", &models.PageData{StrMap: strMap}, request)
}

func (m *Repository) PageHandler(w http.ResponseWriter, request *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, "page.page.tmpl", &models.PageData{StrMap: strMap}, request)
}

func (m *Repository) MakePostHandler(w http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(w, "make-post.page.tmpl", &models.PageData{
		Form: forms.New(nil),
		/*StrMap: strMap*/
	}, request)
	// strMap := make(map[string]string)
}

// Handler for posting articles using post
func (m *Repository) PostMakePostHandler(w http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	article := models.Article{
		BlogTitle:   request.Form.Get("blog_title"),
		BlogArticle: request.Form.Get("blog_acticle"),
	}
	form := forms.New(request.Form)
	form.HasValue("blog_title", request)
	if !form.Valid() {
		data := make(map[string]interface{})
		data["article"] = article
		render.RenderTemplate(w, "make-post.page.tmpl", &models.PageData{
			Form: form,
			Data: data,
		}, request)
		return
	}

	// blog_title := request.Form.Get("blog_title")
	// blog_article := request.Form.Get("blog_acticle")
	// w.Write([]byte(blog_title))
	// w.Write([]byte(blog_article))
	// strMap := make(map[string]string)
	// render.RenderTemplate(w, "page.page.tmpl", &models.PageData{StrMap: strMap})
}
