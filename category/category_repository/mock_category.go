package category_repository

import (
	"database/sql"

	"github.com/FundStation/category"

	//"errors"

	"github.com/FundStation/models"
)

type MockCategoryRepository struct {
	conn *sql.DB
}

// NewPsqlCategoryRepository will create an object of PsqlCategoryRepository
func NewMockCategoryRepository(Conn *sql.DB) category.CategoryRepository {
	return &MockCategoryRepository{conn: Conn}
}

func (cr *MockCategoryRepository) SelectCategory(categoryType string) (categoryy []models.Category, err error) {


		categoryy = append(categoryy,models.CategoryMock)
		return categoryy,nil



}
func (cr *MockCategoryRepository) SelectSpecificCategory(categoryName string) (cat models.Category, err error) {
	cat = models.CategoryMock
	return cat,err

}

