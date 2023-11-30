package handlers

import (
	"net/http"
	"web3/models"
	"web3/pkg/config"
	"web3/pkg/render"
)

type Repository struct{
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository {
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, request *http.Request) {
	m.App.Session.Put(request.Context(), "userid", "sinalalebakhsh") // derek banas
	render.RenderTemplate(w, "home.page.tmpl", &models.PageData{}, request)
}

func (m *Repository) AboutHandler(w http.ResponseWriter, request *http.Request) {
	strMap := make(map[string]string)
	// strMap["title"] = "About Sina Laleh.Bakhsh"
	// strMap["intro"] = "This page is where we talk about ourselves. We love talking about ourselves."
	// userid := m.App.Session.GetString(request.Context(), "userid")
	// strMap["userid"] = userid
	render.RenderTemplate(w, "about.page.tmpl", &models.PageData{StrMap: strMap}, request)
}

func (m *Repository) LoginHandler(w http.ResponseWriter, request *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, "login.page.tmpl", &models.PageData{StrMap: strMap}, request)
}


func (m *Repository) MakePostHandler(w http.ResponseWriter, request *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, "make-post.page.tmpl", &models.PageData{StrMap: strMap}, request)
}


func (m *Repository) PageHandler(w http.ResponseWriter, request *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, "page.page.tmpl", &models.PageData{StrMap: strMap}, request)
}

func (m *Repository) PostMakePostHandler(w http.ResponseWriter, request *http.Request) {
	// strMap := make(map[string]string)
	// render.RenderTemplate(w, "page.page.tmpl", &models.PageData{StrMap: strMap})
	blog_title := request.Form.Get("blog_title")
	blog_article := request.Form.Get("blog_acticle")
	w.Write([]byte(blog_title))
	w.Write([]byte(blog_article))
}






