package handlers

import (
	"Go_Vacay/internal/forms"
	"Go_Vacay/internal/helpers"
	"Go_Vacay/internal/models"
	"Go_Vacay/internal/renderers"
	"fmt"
	"net/http"
)

// Login Page Handler : Get Method
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	var emptyLogin models.Login
	data := make(map[string]interface{})
	data["register"] = emptyLogin

	renderers.RenderTemplate(w, "login.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	}, r)
}

// Login Page Handler
func (m *Repository) PostLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Print("inside postlogin function")
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	register := models.Login{
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}
	form := forms.New(r.PostForm)
	form.Required("email", "password")
	form.IsEmail("email")
	form.IsAlphaNumeric("password")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["register"] = register
		renderers.RenderTemplate(w, "login.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		}, r)
		return
	}
	m.App.Session.Put(r.Context(), "register", register)
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
