package main

import (
	"net/http"

	"Go_Vacay/internal/config"
	"Go_Vacay/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(_ *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(Nosurf)
	mux.Use(SessionLoadandSave)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/login", handlers.Repo.Login)
	mux.Get("/register", handlers.Repo.Register)
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostSearchAvailability)
	mux.Post("/search-availability-json", handlers.Repo.SearchAvailabilityJSON)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Get("/executive-room", handlers.Repo.ExecutiveSuite)
	mux.Get("/deluxe-room", handlers.Repo.Deluxe)
	mux.Get("/premier-room", handlers.Repo.Premier)

	path := config.GetDirPath() + "static"
	fileServer := http.FileServer(http.Dir(path))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}
