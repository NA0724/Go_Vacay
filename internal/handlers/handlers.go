package handlers

import (
	"Go_Vacay/internal/config"
	"Go_Vacay/internal/forms"
	"Go_Vacay/internal/models"
	"Go_Vacay/internal/renderers"
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
	renderers.RenderTemplate(w, "contact.page.html", &models.TemplateData{}, r)
}

// Login Page Handler
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "login.page.html", &models.TemplateData{}, r)
}

// Register/ Sign Up Page Handler
func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "register.page.html", &models.TemplateData{}, r)
}

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

// /Search Availability Page Handler
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "search-availability.page.html", &models.TemplateData{}, r)
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
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	renderers.RenderTemplate(w, "make-reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	}, r)
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		renderers.RenderTemplate(w, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		}, r)
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// ReservationSummary displays the res summary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("can't get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	renderers.RenderTemplate(w, "reservation-summary.page.html", &models.TemplateData{
		Data: data,
	}, r)
}
