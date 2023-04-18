package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateCache    map[string]*template.Template
	Session          *scs.SessionManager
	SecureConnection bool
}
