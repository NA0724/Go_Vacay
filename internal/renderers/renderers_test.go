package renderers

import (
	"Go_Vacay/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "success")
	session.Put(r.Context(), "warning", "warning")
	session.Put(r.Context(), "error", "error")

	result := AddDefaultData(&td, r)
	if result.Flash != "success" {
		t.Error("failed to add default data")
	}
}

// failing because of path
func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "internal/templates"
	templateCache, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = templateCache
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww MyWriter
	err = RenderTemplate(&ww, "home.page.html", &models.TemplateData{}, r)
	if err != nil {
		t.Error("error rendering template to browser")
	}
	err = RenderTemplate(&ww, "house.page.html", &models.TemplateData{}, r)
	if err != nil {
		t.Error("rendered template that does not exist")
	}

}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context() // create a new request context: write to or read from session
	// put session in context
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil

}
