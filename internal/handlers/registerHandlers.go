package handlers

import (
	"Go_Vacay/internal/forms"
	"Go_Vacay/internal/models"
	"Go_Vacay/internal/renderers"
	"log"
	"net/http"
)

// Register/ Sign Up Page Handler : Get Method
func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {
	var emptyRegistration models.Registration
	data := make(map[string]interface{})
	data["register"] = emptyRegistration

	renderers.RenderTemplate(w, "register.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	}, r)
}

// PostRegister page
func (m *Repository) PostRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	register := models.Registration{
		FirstName: r.Form.Get("firstname"),
		LastName:  r.Form.Get("lastname"),
		Email:     r.Form.Get("email"),
		Password:  r.Form.Get("password"),
	}
	form := forms.New(r.PostForm)
	form.Required("firstname", "lastname", "email", "password")
	form.MinLength("firstname", 3)
	form.MinLength("lastname", 1)
	form.MaxLength("firstname", 25)
	form.MaxLength("lastname", 25)
	form.MaxLength("password", 10)
	form.IsEmail("email")
	form.IsAlphaNumeric("password")
	form.IsSame("email", "email1")
	form.IsSame("password", "pwd")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["register"] = register
		renderers.RenderTemplate(w, "register.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		}, r)
		return
	}

	m.App.Session.Put(r.Context(), "register", register)
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

// My profile page
func (m *Repository) MyProfile(w http.ResponseWriter, r *http.Request) {
	register, ok := m.App.Session.Get(r.Context(), "register").(models.Registration)
	if !ok {
		log.Println("Can't get user item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get user from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "register")

	data := make(map[string]interface{})
	data["register"] = register

	renderers.RenderTemplate(w, "myprofile.page.html", &models.TemplateData{
		Data: data,
	}, r)
}
