package handlers

import (
	"log"
	"net/http"
	"strconv"
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
	if m.App.Session.Get(r.Context(), "current-district") == nil {
		m.App.Session.Put(r.Context(), "current-district", districts[0])
	}

	data := make(map[string]interface{})
	data["districts"] = districts
	data["current-district"] = m.App.Session.Get(r.Context(), "current-district")

	err := render.Template(w, r, "home.page.tmpl", &models.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	districts := []models.District{
		{1, "Мурино", time.Now(), time.Now()},
		{2, "Колтуши", time.Now(), time.Now()},
	}
	if m.App.Session.Get(r.Context(), "current-district") == nil {
		m.App.Session.Put(r.Context(), "current-district", districts[0])
	}

	data := make(map[string]interface{})
	data["districts"] = districts
	data["current-district"] = m.App.Session.Get(r.Context(), "current-district")

	err := render.Template(w, r, "about.page.tmpl", &models.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) Districts(w http.ResponseWriter, r *http.Request) {
	districts := []models.District{
		{1, "Мурино", time.Now(), time.Now()},
		{2, "Колтуши", time.Now(), time.Now()},
	}
	if m.App.Session.Get(r.Context(), "current-district") == nil {
		m.App.Session.Put(r.Context(), "current-district", districts[0])
	}

	data := make(map[string]interface{})
	data["districts"] = districts
	data["current-district"] = m.App.Session.Get(r.Context(), "current-district")

	err := render.Template(w, r, "districts.page.tmpl", &models.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) ChangeDistrict(w http.ResponseWriter, r *http.Request) {
	m.App.Session.Remove(r.Context(), "current-district")
	id, _ := strconv.Atoi(r.URL.Query()["id"][0])
	name := r.URL.Query()["name"][0]

	m.App.Session.Put(r.Context(), "current-district", models.District{ID: id, Name: name})

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
