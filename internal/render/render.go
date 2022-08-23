package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sushipizza/internal/config"
	"sushipizza/internal/models"
)

var app *config.AppConfig
var pathToTemplates = "./templates"
var functions template.FuncMap

func NewRenderer(a *config.AppConfig) {
	app = a
}

// CreateTemplateCache creates template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		log.Println(err)
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println(err)
			return myCache, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			log.Println(err)
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				log.Println(err)
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

// Template renders page
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	log.Println(ok)
	if !ok {
		log.Println("cannot get template from cache")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("buf len", buf.Len())
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
		return err
	}

	return nil
}
