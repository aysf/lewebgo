package handlers

import (
	"net/http"

	"github.com/aysf/lewebgo/config"
	"github.com/aysf/lewebgo/pkg/model"
	"github.com/aysf/lewebgo/pkg/render"
)

var Repo *Repository

type Repository struct {
	app config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{app: *a}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	m.app.Session.Put(r.Context(), "ipAddr", r.RemoteAddr)

	render.RenderTemplate(w, "home.page.html", &model.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	sm := map[string]string{}
	sm["greeting"] = "Hello Template Data"
	sm["ip"] = m.app.Session.GetString(r.Context(), "ipAddr")

	render.RenderTemplate(w, "about.page.html", &model.TemplateData{
		StringMap: sm,
	})
}
