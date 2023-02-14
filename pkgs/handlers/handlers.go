package handlers

import (
	"Go_Vacay/pkgs/config"
	"Go_Vacay/pkgs/models"
	"Go_Vacay/pkgs/renderers"
	"encoding/json"
	"fmt"
	"log"
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
	renderers.RenderTemplate(w, "contact.html", &models.TemplateData{}, r)
}

// Login Page Handler
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "login.html", &models.TemplateData{}, r)
}

// Register/ Sign Up Page Handler
func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "register.html", &models.TemplateData{}, r)
}

// Executive Page Handler
func (m *Repository) ExecutiveSuite(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "executive-room.html", &models.TemplateData{}, r)
}

// Deluxe Page Handler
func (m *Repository) Deluxe(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "deluxe-room.html", &models.TemplateData{}, r)
}

// /Premier Page Handler
func (m *Repository) Premier(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "premier-room.html", &models.TemplateData{}, r)
}

// /Search Availability Page Handler
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "search-availability.html", &models.TemplateData{}, r)
}

// Post Search Availability Page Handler
func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("startdate")
	end := r.Form.Get("enddate")
	w.Write([]byte(fmt.Sprintf("start date %s and end date %s", start, end)))

}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"msg"`
}

// Search Availability JSON Page Handler
func (m *Repository) SearchAvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	respone := jsonResponse{
		OK:      true,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(respone, "", " ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Make Reservation Handler
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "make-reservation.html", &models.TemplateData{}, r)
}
