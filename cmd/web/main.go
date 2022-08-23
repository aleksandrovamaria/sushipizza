package main

import (
	"log"
	"net/http"
	"sushipizza/internal/config"
	"sushipizza/internal/handlers"
	"sushipizza/internal/render"
)

const port = ":80"

var app config.AppConfig

func main() {
	app.InProduction = false
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)

	log.Println("Starting an application on port", port)
	srv := http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
