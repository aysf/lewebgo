package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func MyCustomlogger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("hit the page: %v\n", r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfhandler := nosurf.New(next)

	csrfhandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfhandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
