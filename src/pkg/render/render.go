package render

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/saulaverageman/ugly-go-bnb/pkg/config"
)

var appConfig *config.AppConfig
var err error

func NewRender(config *config.AppConfig) {
	appConfig = config
}

// renders and writes templats based on given tmpl name
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	WORKPATH := "/work/"

	// getting the templateCache from the cental app config
	templateCache := appConfig.TemplateCache

	// loading the file from cache
	parsedTemplate, fileExists := templateCache[tmpl]
	if !fileExists {
		log.Panic("error: template cache is not available, loading from disk")

		// loading the specific file into cache as the file is not found in the cache
		templateCache[tmpl], err = template.ParseFiles(WORKPATH+"templates/"+tmpl, WORKPATH+"templates/base.layout.tmpl")

		if err != nil {
			log.Fatal("error: loading", tmpl, "beacuse", err)
			return
		}
	}

	// writing the file in http response writer
	log.Println("verbose: serving from cache template-", tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Fatal("error rendering template: ", tmpl, "with error:", err)
		return
	}
}

// func to create a cache with all the templates
func FormTemplateCache(WORKPATH string) (map[string]*template.Template, error) {
	log.Print("verbose: loading template cache")

	tc := map[string]*template.Template{}

	files, err := os.ReadDir(WORKPATH + "templates/")
	if err != nil {
		log.Fatal("error: cannot read files at WORKPATH,", err)
		return nil, err
	}

	for _, file := range files {
		tc[file.Name()], err = template.ParseFiles(WORKPATH+"templates/"+file.Name(), WORKPATH+"templates/base.layout.tmpl")
		if err != nil {
			log.Panic("error: cannot read file", WORKPATH+"templates/"+file.Name())
			return nil, err
		}
	}
	return tc, nil
}
