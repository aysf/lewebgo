package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

var port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is Home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is About Page")
}

func Divide(w http.ResponseWriter, r *http.Request) {
	res, err := dividesValue(4, 0)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Fprintf(w, "the result is %f", res)
	}
}

func dividesValue(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	return x / y, nil
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Web")
	})

	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/div", Divide)

	log.Println("server running on port: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("error running server: ", err)
	}
}
