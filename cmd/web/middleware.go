package main

import (
	"fmt"
	"net/http"
)

func MyCustomlogger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("hit the page: %v\n", r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
