package category

import "github.com/FundStation/models"

type CategoryService interface {
	ViewCategory(string) ([]models.Category, error)
	ViewSpecificCategory(string) (models.Category, error)
}