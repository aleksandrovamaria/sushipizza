package repository

import "sushipizza/internal/models"

type DatabaseRepo interface {
	AllDistricts() []models.District
	AllCategories() []models.Category
	AllItemsOfCategory(id int) []models.Item
}
