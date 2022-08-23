package main

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"sushipizza/internal/config"
	"sushipizza/internal/handlers"
	"sushipizza/internal/models"
	"sushipizza/internal/render"
	"time"
)

const port = ":80"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false
	app.UseCache = app.InProduction

	session = scs.New()
	session.Lifetime = 30 * time.Minute
	session.Cookie.Secure = app.InProduction
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	gob.Register(models.District{})

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
