package handler

import (
	"net/http"

	"github.com/saulaverageman/ugly-go-bnb/pkg/config"
	"github.com/saulaverageman/ugly-go-bnb/pkg/render"
)

var appConfig *config.AppConfig

//getting app config
func NewHandler(config *config.AppConfig) {
	appConfig = config
}

// render home tmpl
func Home(w http.ResponseWriter, r *http.Request) {
	appConfig.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)
	render.RenderTemplate(w, "home.tmpl")
}

// renders about tmpl
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.tmpl")
}
