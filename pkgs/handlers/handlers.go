package handlers

import (
	"net/http"

	"github.com/NA0724/Go_Vacay/pkgs/config"
	"github.com/NA0724/Go_Vacay/pkgs/models"
	"github.com/NA0724/Go_Vacay/pkgs/renderers"
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

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "Remote_IP", remoteIP) // storing the remote ip in the session everytime the request to homepage is made
	renderers.RenderTemplate(w, "homepage.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"
	Remote_IP := m.App.Session.GetString(r.Context(), "Remote_IP")
	stringMap["remote_ip"] = Remote_IP
	renderers.RenderTemplate(w, "aboutpage.html", &models.TemplateData{StringMap: stringMap})
}
