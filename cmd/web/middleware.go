package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// adds CSRF protection to all POST requests
func Nosurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProd,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Loads and Saves the session on every request
func SessionLoadandSave(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
