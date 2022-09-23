package main

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, t string) {
	parsedTemplate, err := template.ParseFiles(t)
	if err != nil {
		log.Panic("error parsing files", err)
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Panic("error execute template", err)
	}
}
