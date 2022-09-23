package handlers

import (
	"net/http"

	"github.com/aysf/lewebgo/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "./templates/home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "./templates/about.html")
}
