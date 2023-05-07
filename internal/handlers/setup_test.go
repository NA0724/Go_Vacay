package handlers

import (
	"Go_Vacay/internal/config"
	"Go_Vacay/internal/models"
	"Go_Vacay/internal/renderers"
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager

func getRoutes() http.Handler {
	gob.Register(models.Reservation{})
	gob.Register(models.Registration{})
	//set to true if production environment
	app.InProd = false

	//initialise session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd

	app.Session = session

	tempCache, err := renderers.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tempCache
	app.UseCache = false // false for development mode, true for prod or qa mode

	repo := NewRepo(&app)
	NewHandlers(repo)

	renderers.NewTemplates(&app)
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(Nosurf)
	mux.Use(SessionLoadandSave)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)

	mux.Get("/login", Repo.Login)
	mux.Post("/login", Repo.PostLogin)

	mux.Get("/register", Repo.Register)
	mux.Post("/register", Repo.PostRegister)

	mux.Get("/profile", Repo.MyProfile)

	mux.Get("/search-availability", Repo.SearchAvailability)
	mux.Post("/search-availability", Repo.PostSearchAvailability)
	mux.Post("/search-availability-json", Repo.SearchAvailabilityJSON)

	mux.Get("/contact", Repo.Contact)
	mux.Post("/contact", Repo.Contact)
	mux.Get("/success", Repo.Contact)

	mux.Get("/make-reservation", Repo.MakeReservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	mux.Get("/executive-room", Repo.ExecutiveSuite)
	mux.Get("/deluxe-room", Repo.Deluxe)
	mux.Get("/premier-room", Repo.Premier)

	path := config.GetDirPath() + "static"
	fileServer := http.FileServer(http.Dir(path))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

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
