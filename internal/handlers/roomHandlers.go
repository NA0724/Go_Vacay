package handlers

import (
	"Go_Vacay/internal/models"
	"Go_Vacay/internal/renderers"
	"net/http"
)

// Executive Page Handler
func (m *Repository) ExecutiveSuite(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "executive-room.page.html", &models.TemplateData{}, r)
}

// Deluxe Page Handler
func (m *Repository) Deluxe(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "deluxe-room.page.html", &models.TemplateData{}, r)
}

// /Premier Page Handler
func (m *Repository) Premier(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "premier-room.page.html", &models.TemplateData{}, r)
}
