package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/aysf/lewebgo/config"
	"github.com/aysf/lewebgo/pkg/handlers"
	"github.com/aysf/lewebgo/pkg/render"
)

var port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	var err error

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	// this setting will persist the cookie after user close the browser
	session.Cookie.Persist = true
	// how strict you wanna be able waht cookie applied to
	session.Cookie.SameSite = http.SameSiteLaxMode
	// http or https
	session.Cookie.Secure = app.InProduction

	app.Session = session

	app.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache:", err)
	}

	app.UseCache = true

	render.NewTemplates(&app)
	r := handlers.NewRepo(&app)
	handlers.NewHandler(r)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Web")
	})

	// http.HandleFunc("/home", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	log.Println("server running on port", port)
	err = http.ListenAndServe(port, Route(&app))
	if err != nil {
		log.Println("error running server: ", err)
	}
}
