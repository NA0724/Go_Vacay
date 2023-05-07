package renderers

import (
	"Go_Vacay/internal/config"
	"Go_Vacay/internal/models"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
)

var app *config.AppConfig
var pathToTemplates = "internal/templates/"

// sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds data for all templates
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, tempdata *models.TemplateData, r *http.Request) {

	var tc map[string]*template.Template
	var er error
	if app.UseCache {
		// if UseCache is true then fetch template from the cache, get template cache from the app config
		tc = app.TemplateCache
	} else {
		// Else, build cache from scratch and load from disk
		tc, er = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache", er)
	}
	buffer := new(bytes.Buffer)
	tempdata = AddDefaultData(tempdata, r) //default data
	_ = t.Execute(buffer, tempdata)        //pass the data here
	_, err := buffer.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

// create a template cache and return the map,
// automatically acesses all files in the templates directory, no need to add manually
func CreateTemplateCache() (map[string]*template.Template, error) {

	tempCache := map[string]*template.Template{}

	//get all files named *.page.html from ./templates
	pages, err := filepath.Glob(config.GetDirPath() + fmt.Sprintf("%s/*.page.html", pathToTemplates))

	if err != nil {
		return tempCache, err
	}

	//range through all pages ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)                    // get only the last element in path, i.e name of the file
		ts, err := template.New(name).ParseFiles(page) // ts pointer to template
		if err != nil {
			return tempCache, err
		}
		// find the layout template
		matches, er := filepath.Glob(config.GetDirPath() + "internal/templates/*.layout.html") // find layout file
		if er != nil {
			return tempCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(config.GetDirPath() + "internal/templates/*.layout.html") // parse layout file
			if err != nil {
				return tempCache, err
			}
		}
		tempCache[name] = ts
	}
	return tempCache, nil
}
