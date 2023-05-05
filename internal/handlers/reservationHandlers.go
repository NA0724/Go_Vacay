package handlers

import (
	"Go_Vacay/internal/forms"
	"Go_Vacay/internal/models"
	"Go_Vacay/internal/renderers"
	"log"
	"net/http"
)

// Make Reservation Handler : Get mathod
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
		FirstName: r.Form.Get("firstname"),
		LastName:  r.Form.Get("lastname"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		Checkin:   r.Form.Get("startdate"),
		Checkout:  r.Form.Get("enddate"),
	}

	form := forms.New(r.PostForm)

	form.Required("firstname", "lastname", "email", "phone")
	form.MinLength("firstname", 3, r)
	form.MinLength("lastname", 1, r)
	form.MaxLength("firstname", 25, r)
	form.MaxLength("lastname", 25, r)
	form.MaxLength("phone", 10, r)
	form.Is_numeric("phone", r)
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
		log.Println("Can't get item from session")
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
