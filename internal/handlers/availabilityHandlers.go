package handlers

import (
	"Go_Vacay/internal/helpers"
	"Go_Vacay/internal/models"
	"Go_Vacay/internal/renderers"
	"encoding/json"
	"fmt"
	"net/http"
)

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
	out, err := json.MarshalIndent(respone, "", " 	")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
