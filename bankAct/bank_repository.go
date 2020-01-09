package bankAct

import (
	"github.com/FundStation/models"
)

type BankRepository interface {
	SelectAccountNo(string) (models.BankAccount, error)
	UpdateBalance(string, float64) error
}
