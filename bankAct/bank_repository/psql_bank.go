package bank_repository

import (
	"database/sql"
	"errors"
	"github.com/FundStation/models"
	_ "github.com/lib/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type PsqlBankRepository struct {
	conn *sql.DB
}

func NewPsqlBankRepository(Conn *sql.DB) *PsqlBankRepository {
	return &PsqlBankRepository{conn: Conn}
}

func (pr *PsqlBankRepository) SelectAccountNo(aNo string) (act models.BankAccount, err error) {
	err = pr.conn.QueryRow("SELECT AccountNo, Balance  FROM bankInfo WHERE AccountNo=$1", aNo).Scan(&act.AccountNo, &act.CurrentBalance)
	if err != nil {
		return act, errors.New("Incorrect AccountNo")
	}

	return act, nil
}
func (pr *PsqlBankRepository) UpdateBalance(aNo string, newBalnc float64) error {

	_, err := pr.conn.Exec("UPDATE bankInfo SET Balance=$1 WHERE AccountNo =$2", newBalnc, aNo)

	if err != nil {
		return errors.New("Could not update")

	}
	return nil

}
