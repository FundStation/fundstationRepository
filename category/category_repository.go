package category

import "github.com/FundStation/models"

type CategoryRepository interface {
	SelectCategory(string) ([]models.Category, error)
	SelectSpecificCategory(string)(models.Category,error)
}
