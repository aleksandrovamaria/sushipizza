package handlers

import (
	"log"
	"net/http"
	"sushipizza/internal/config"
	"sushipizza/internal/models"
	"sushipizza/internal/render"
	"time"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	districts := []models.District{
		{1, "Мурино", time.Now(), time.Now()},
		{2, "Колтуши", time.Now(), time.Now()},
	}
	data := make(map[string]interface{})
	data["districts"] = districts
	err := render.Template(w, r, "home.page.tmpl", &models.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) Districts(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "districts.page.tmpl", &models.TemplateData{})
	if err != nil {
		log.Println(err)
	}
}
