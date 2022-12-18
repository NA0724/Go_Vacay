package renderers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/NA0724/Go_Vacay/pkgs/config"
	"github.com/NA0724/Go_Vacay/pkgs/models"
)

var app *config.AppConfig

// sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, tempdata *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		// if UseCache is true then fetch template from the cache, get template cache from the app config
		tc = app.TemplateCache
	} else {
		// Else, build cache from scratch and load from disk
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}
	buffer := new(bytes.Buffer)
	tempdata = AddDefaultData(tempdata) //default data
	_ = t.Execute(buffer, tempdata)     //pass the data here
	_, err := buffer.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

// create a template cache and return the map,
// automatically acesses all files in the templates directory, no need to add manually
func CreateTemplateCache() (map[string]*template.Template, error) {

	tempCache := map[string]*template.Template{}

	//get all files named *page.html from ./templates
	pages, err := filepath.Glob("./pkgs/templates/*page.html")
	if err != nil {
		return tempCache, err
	}

	//range through all pages ending with *page.html
	for _, page := range pages {
		name := filepath.Base(page)                    // get only the last element in path, i.e name of the file
		ts, err := template.New(name).ParseFiles(page) // ts pointer to template
		if err != nil {
			return tempCache, err
		}

		// find the layout template
		matches, er := filepath.Glob("./pkgs/templates/*.layout.html") // find layout file
		if er != nil {
			return tempCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./pkgs/templates/*.layout.html") // parse layout file
			if err != nil {
				return tempCache, err
			}
		}
		tempCache[name] = ts
	}
	return tempCache, nil
}
