package category_repository

import (
	"database/sql"
	"errors"

	//"errors"

	"github.com/FundStation/models"
)

type PsqlCategoryRepository struct {
	conn *sql.DB
}

// NewPsqlCategoryRepository will create an object of PsqlCategoryRepository
func NewPsqlCategoryRepository(Conn *sql.DB) *PsqlCategoryRepository {
	return &PsqlCategoryRepository{conn: Conn}
}

func (pr *PsqlCategoryRepository) SelectCategory(categoryType string) (categoryy []models.Category, err error) {
	selected, err := pr.conn.Query("SELECT namee, image FROM category WHERE typee=$1", categoryType)
	if err != nil {
		return categoryy, errors.New("something")
	}

	cat := models.Category{}
	for selected.Next() {
		err := selected.Scan(&cat.Name, &cat.Image)

		if err != nil {
			return categoryy, errors.New("Couldnot")
		}
		categoryy = append(categoryy, cat)
	}
	return categoryy, nil


}
func (pr *PsqlCategoryRepository) SelectSpecificCategory(categoryName string) (cat models.Category, err error) {
	selected,err:= pr.conn.Prepare("SELECT namee, description, account, image FROM category WHERE namee=$1")
	if err != nil {
		return cat, errors.New("something")
	}
	var bankAccount models.BankAccount
	cat= models.Category{}

		err = selected.QueryRow(categoryName).Scan(&cat.Name, &cat.Description, &bankAccount.AccountNo, &cat.Image)
		cat.BankAccount = bankAccount
		if err != nil {
			return cat, errors.New("Couldnot")

	}
	return cat, nil


}
