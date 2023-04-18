package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/saulaverageman/ugly-go-bnb/pkg/config"
	"github.com/saulaverageman/ugly-go-bnb/pkg/handler"
)

var appConfig *config.AppConfig

func routes(app *config.AppConfig) http.Handler {
	appConfig = app

	mux := chi.NewMux()

	//Middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad) //Loading session cookie as server by nature doesnt do on its own

	//routes
	handler.NewHandler(appConfig)
	mux.Get("/", handler.Home)
	mux.Get("/about", handler.About)

	return mux
}
