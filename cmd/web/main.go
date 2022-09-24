package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aysf/lewebgo/config"
	"github.com/aysf/lewebgo/pkg/handlers"
	"github.com/aysf/lewebgo/pkg/render"
)

var port = ":8080"

func main() {

	var app config.AppConfig
	var err error

	app.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache:", err)
	}

	render.NewTemplates(&app)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Web")
	})

	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("server running on port", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("error running server: ", err)
	}
}
