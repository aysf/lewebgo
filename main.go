package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = ":8080"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Web")
	})

	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)

	log.Println("server running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("error running server: ", err)
	}
}
