package bank_repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/FundStation/models"
	//_ "github.com/lib/go-sql-driver/mysql"
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
func (pr *PsqlBankRepository) AccountExists(account string) bool {
	var balance string
	err := pr.conn.QueryRow("SELECT balance FROM bankInfo WHERE accountno=$1",account).Scan(&balance)

	if err != nil {
		fmt.Println("accountExist",account)
		fmt.Println("bankerr",err)
		return true
	}

	return false
}
