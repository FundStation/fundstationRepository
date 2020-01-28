package disaster_repository

import (
	"database/sql"
	//"errors"
	"github.com/FundStation/models"
)

type PsqlDisasterRepository struct {
	conn *sql.DB
}

func NewPsqlDisasterRepository(Conn *sql.DB) *PsqlDisasterRepository {
	return &PsqlDisasterRepository{conn: Conn}
}

func (pr *PsqlDisasterRepository) SelectDisaster() (disaster []models.Disaster, err error) {
	selected, err := pr.conn.Query("SELECT namee, description, account, image FROM disaster")
	if err != nil {
		panic(err)
	}

	dis := models.Disaster{}
	var bankAccount models.BankAccount
	for selected.Next() {

		err := selected.Scan(&dis.Name,&dis.Description, &bankAccount.AccountNo, &dis.Image)
		dis.BankAccount = bankAccount
		if err != nil {
			panic(err)
		}
		disaster = append(disaster, dis)
	}
	return disaster, nil

}
func (pr *PsqlDisasterRepository) SelectEvent() (event []models.Event, err error) {
	selected, err := pr.conn.Query("SELECT namee, description, image FROM event")
	if err != nil {
		panic(err)
	}

	eve := models.Event{}
	//var bankAccount models.BankAccount
	for selected.Next() {
		err := selected.Scan(&eve.Name,&eve.Description, &eve.Image)
		//dis.BankAccount = bankAccount
		if err != nil {
			panic(err)
		}
		event = append(event, eve)
	}
	return event, nil

}
