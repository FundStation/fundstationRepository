package recipient_repository

import (
	"database/sql"
	"errors"
	"github.com/FundStation/models"
)


type PsqlRecipientRepository struct {
	conn *sql.DB
}

func NewPsqlRecipientRepository(Conn *sql.DB) *PsqlRecipientRepository {
	return &PsqlRecipientRepository{conn: Conn}
}

func (pr *PsqlRecipientRepository) InsertRecipient(r models.Recipient) error {

	err := pr.conn.QueryRow("INSERT INTO recipient(firstname,lastname,address,username,password,phone,email) VALUES($1, $2, $3, $4, $5, $6, $7) returning id", r.FirstName, r.LastName, r.Address, r.Username, r.Password, r.PhoneNumber, r.EmailAddress).Scan(&r.RecipientNo)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PsqlRecipientRepository) SelectRecipient(r models.Recipient) error {

	querystmt, err := pr.conn.Prepare("SELECT username,password FROM recipient WHERE username=$1 AND password=$2")

	if err != nil {
		return err
	}
	var username string
	var pass string

	err = querystmt.QueryRow(r.Username, r.Password).Scan(&username, &pass)

	if err == sql.ErrNoRows {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (pr *PsqlRecipientRepository) SelectAllRecipient() (recipients []models.Recipient, err error) {
	selRec, err := pr.conn.Query("SELECT id ,firstname, lastname, address,phone,email FROM recipient")
	if err != nil {
		return recipients, errors.New("something")
	}
	recp := models.Recipient{}
	for selRec.Next() {
		err := selRec.Scan(&recp.RecipientNo, &recp.FirstName, &recp.LastName, &recp.Address,  &recp.PhoneNumber, &recp.EmailAddress)
		if err != nil {
			return recipients, errors.New("Could not select recipient")
		}
		recipients = append(recipients, recp)
	}
	return recipients, nil
}

