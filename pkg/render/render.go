package render

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
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

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, templ string) {

// 	var t *template.Template
// 	var err error

// 	_, ok := tc[templ]
// 	if !ok {
// 		// need create template
// 		log.Println("creating template cache")
// 		err = createTemplateCache(templ)
// 		if err != nil {
// 			log.Println("error creating template:", err)
// 		}
// 	} else {
// 		// we have the template in the cache
// 		log.Println("using template cache")
// 	}

// 	t = tc[templ]

// 	err = t.Execute(w, nil)
// 	if err != nil {
// 		log.Fatal("error executing template", err)
// 	}

// }

// func createTemplateCache(t string) error {

// 	parsedTemplate, err := template.ParseFiles("./templates/"+t, "./templates/base.html")
// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = parsedTemplate

// 	return nil
// }

func RenderTemplate(w http.ResponseWriter, templ string) {

	// tMap := make(map[string]*template.Template)
	var err error

	tMap, err := createTemplateCache()
	if err != nil {
		log.Panic("error creating template:", err)
	}

	t, ok := tMap[templ]
	if !ok {
		log.Panic(errors.New("template not found"))
	}
	// t := tMap[templ]

	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal("error executing template", err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	cacheMap := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return cacheMap, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		tc, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cacheMap, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return cacheMap, err
		}

		if len(matches) > 0 {

			tc, err = tc.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return cacheMap, err
			}
		}

		cacheMap[name] = tc

	}

	return cacheMap, nil
}
