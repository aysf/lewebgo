package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aysf/lewebgo/pkg/handlers"
)

var port = ":8080"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Web")
	})

	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("server running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("error running server: ", err)
	}
}
