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
	render.RenderTemplate(w, request, "home.page.tmpl", &models.PageData{})
}

func (m *Repository) AboutHandler(w http.ResponseWriter, request *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, request, "about.page.tmpl", &models.PageData{StrMap: strMap})
	// strMap["title"] = "About Sina Laleh.Bakhsh"
	// strMap["intro"] = "This page is where we talk about ourselves. We love talking about ourselves."
	// userid := m.App.Session.GetString(request.Context(), "userid")
	// strMap["userid"] = userid
}

func (m *Repository) LoginHandler(w http.ResponseWriter, request *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, request, "login.page.tmpl", &models.PageData{StrMap: strMap})
}

func (m *Repository) PageHandler(w http.ResponseWriter, request *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, request, "page.page.tmpl", &models.PageData{StrMap: strMap})
}

func (m *Repository) MakePostHandler(w http.ResponseWriter, request *http.Request) {

	var emptyArticle models.Article
	data := make(map[string]interface{})
	data["article"] = emptyArticle

	render.RenderTemplate(w, request, "make-post.page.tmpl", &models.PageData{
		Form: forms.New(nil),
		Data: data,
		/*StrMap: strMap*/
	})
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
		BlogArticle: request.Form.Get("blog_article"),
	}
	form := forms.New(request.Form)
	// form.HasValue("blog_title", request)
	form.HasRequired("blog_title", "blog_article")

	form.MinLength("blog_title", 5, request)
	form.MinLength("blog_article", 5, request)


	// form.IsEmail("email")
	

	if !form.Valid() {
		data := make(map[string]interface{})
		data["article"] = article
		render.RenderTemplate(w, request, "make-post.page.tmpl", &models.PageData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(request.Context(), "article", article)
	http.Redirect(w, request, "/article-received", http.StatusSeeOther)


	// blog_title := request.Form.Get("blog_title")
	// blog_article := request.Form.Get("blog_acticle")
	// w.Write([]byte(blog_title))
	// w.Write([]byte(blog_article))
	// strMap := make(map[string]string)
	// render.RenderTemplate(w, "page.page.tmpl", &models.PageData{StrMap: strMap})
}


func (m *Repository) ArticleReceived(w http.ResponseWriter, request *http.Request) {
	article, ok := m.App.Session.Get(request.Context(), "article").(models.Article)
	if !ok {
		log.Println("Can't get data from session")

		m.App.Session.Put(request.Context(), "error", "Can't get data from session")

		http.Redirect(w, request, "/", http.StatusTemporaryRedirect)
		
		return
	}
	data := make(map[string]interface{})
	data["article"] = article

	render.RenderTemplate(w, request, "article-received.page.tmpl", &models.PageData{
		Data: data,
	})

}







