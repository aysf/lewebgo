package render

import (
	"html/template"
	"log"
	"net/http"
)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	parsedTemplate, err := template.ParseFiles(t, "./templates/base.html")
// 	if err != nil {
// 		log.Panic("error parsing files", err)
// 	}
// 	err = parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		log.Panic("error execute template", err)
// 	}
// }

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, templ string) {

	var t *template.Template
	var err error

	_, ok := tc[templ]
	if !ok {
		// need create template
		log.Println("creating template cache")
		err = createTemplateCache(templ)
		if err != nil {
			log.Println("error creating template:", err)
		}
	} else {
		// we have the template in the cache
		log.Println("using template cache")
	}

	t = tc[templ]

	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal("error executing template", err)
	}

}

func createTemplateCache(t string) error {

	parsedTemplate, err := template.ParseFiles("./templates/"+t, "./templates/base.html")
	if err != nil {
		return err
	}

	tc[t] = parsedTemplate

	return nil
}
