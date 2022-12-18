package renderers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// template cache
var tc = make(map[string]*template.Template)

// -------------------------------------- METHOD 1 -------------------------------------------------
// layman approach
/*func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, er1 := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	if er1 != nil {
		fmt.Println("Error in -------- ", er1)
	}
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}*/

// -------------------------------------- METHOD 2 -------------------------------------------------
// better way of accessing templates
func RenderTemplateT(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check to see if we alredy have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		//need to create the template
		log.Println("creating template and adding to template cache")
		err = createTemplate(t)
		if err != nil {
			log.Println("error parsing ", err)
		}
	} else {
		//we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("error template execution: ", err)
	}

}

// create a template and store in cache if not present
func createTemplate(t string) error {
	templates := []string{
		fmt.Sprintf("./pkgs/templates/%s", t),
		"./pkgs/templates/base.layout.html",
	}

	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache map
	tc[t] = tmpl
	return nil
}

// -------------------------------------- METHOD 3 -------------------------------------------------
// complex way to automatically populate template cache
func RenderTemplateTT(w http.ResponseWriter, tmpl string) {

	//Step1: create a template cache
	tempCache, err := CreateTemplateCacheT()
	if err != nil {
		log.Fatal(err)
	}

	//Step2: get requested template from cache
	t, ok := tempCache[tmpl] // t is the template value in the map for key tmpl
	if !ok {
		log.Fatal(err)
	}
	buffer := new(bytes.Buffer)  // hold bytes
	err = t.Execute(buffer, nil) // better error checking
	if err != nil {
		log.Println(err) // error comes from the value stored in the map
	}

	//Step3: render the template
	_, err = buffer.WriteTo((w))
	if err != nil {
		log.Println(err)
	}
}

// create a template cache and return the map,
// automatically acesses all files in the templates directory, no need to add manually
func CreateTemplateCacheT() (map[string]*template.Template, error) {

	//myCache := make(map[string]*template.Template) ------ layman way of creating map
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
