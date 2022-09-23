package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "./templates/home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "./templates/about.html")
}

func renderTemplate(w http.ResponseWriter, t string) {
	parsedTemplate, err := template.ParseFiles(t)
	if err != nil {
		log.Panic("error parsing files", err)
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Panic("error execute template", err)
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Web")
	})

	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)

	log.Println("server running on port: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("error running server: ", err)
	}
}
