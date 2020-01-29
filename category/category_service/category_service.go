package category_service

import (
	"github.com/FundStation/category"
	"github.com/FundStation/models"
)

type CategoryService struct {
	cRepo category.CategoryRepository
}


func NewCategoryService(catRepo category.CategoryRepository) *CategoryService {
	return &CategoryService{cRepo: catRepo}
}

func (cs *CategoryService) ViewCategory(typee string) (categoryy []models.Category, err error) {

	categoryy, err = cs.cRepo.SelectCategory(typee)

	if err != nil {
		return categoryy, err
	}

	return categoryy, nil
}
func (cs *CategoryService) ViewSpecificCategory(catName string) (categoryy models.Category, err error) {

	categoryy, err = cs.cRepo.SelectSpecificCategory(catName)

	if err != nil {
		return categoryy, err
	}

	return categoryy, nil
}
