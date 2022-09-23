package main

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "./templates/home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "./templates/about.html")
}
