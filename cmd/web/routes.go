package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"sushipizza/internal/config"
	"sushipizza/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/districts", handlers.Repo.Districts)
	mux.Get("/change-district", handlers.Repo.ChangeDistrict)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
