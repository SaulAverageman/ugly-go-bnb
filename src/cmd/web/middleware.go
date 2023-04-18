package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

/*
Middleware blueprint

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("a page is hit")
		next.ServeHTTP(w, r)
	})
}
*/

// setting up cross site repliction forgery token protection for all POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.SecureConnection,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// making the server stateful by loading and saving session for every new request
func SessionLoad(next http.Handler) http.Handler {
	return appConfig.Session.LoadAndSave(next)
}
