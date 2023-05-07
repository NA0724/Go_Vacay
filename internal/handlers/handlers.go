package handlers

import (
	"Go_Vacay/internal/config"
	"Go_Vacay/internal/models"
	"Go_Vacay/internal/renderers"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the reporsitory for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home Page Handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "home.page.html", &models.TemplateData{}, r)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "about.page.html", &models.TemplateData{}, r)
}

// Contact Page Handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "contact.page.html", &models.TemplateData{}, r)
}
