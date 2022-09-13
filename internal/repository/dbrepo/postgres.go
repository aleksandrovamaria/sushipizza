package dbrepo

import (
	"sushipizza/internal/models"
)

func (m *postgresDBRepo) AllDistricts() []models.District {
	return nil
}

func (m *postgresDBRepo) AllCategories() []models.Category {
	return nil
}

func (m *postgresDBRepo) AllItemsOfCategory(id int) []models.Item {
	return nil
}
