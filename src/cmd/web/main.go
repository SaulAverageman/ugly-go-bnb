package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/saulaverageman/ugly-go-bnb/pkg/config"
	"github.com/saulaverageman/ugly-go-bnb/pkg/render"
)

// app is the app config for the whole project
var app config.AppConfig

func main() {
	// set true in production
	app.SecureConnection = false

	// session
	session := scs.New()
	session.Lifetime = 12 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = session

	// sending app config to render.go
	render.NewRender(&app)

	// render template cache as soon as the app starts and saving them in appConfig.TemplateCache
	tc, err := render.FormTemplateCache("/work/")
	if err != nil {
		log.Fatal("error: cannot create template cache")
	} else {
		app.TemplateCache = tc
	}

	const appPort = ":8000"

	server := &http.Server{
		Addr:    appPort,
		Handler: routes(&app),
	}

	log.Println("verbose: starting servet @port:", appPort)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
