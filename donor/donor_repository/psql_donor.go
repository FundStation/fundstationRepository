package donor_repository

import (
	"database/sql"
	"errors"

	//"errors"

	"github.com/FundStation/models"
)

type PsqlDonorRepository struct {
	conn *sql.DB
}

func NewPsqlDonorRepository(Conn *sql.DB) *PsqlDonorRepository {
	return &PsqlDonorRepository{conn: Conn}
}

func (pr *PsqlDonorRepository) InsertDonor(d *models.Donor) error {

	err := pr.conn.QueryRow("INSERT INTO donor(firstname,lastname,address,username,password,phone,email) VALUES($1, $2, $3, $4, $5, $6, $7) returning id", d.FirstName, d.LastName, d.Address, d.Username, d.Password, d.PhoneNumber, d.EmailAddress).Scan(&d.DonorNo)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PsqlDonorRepository) SelectDonor(d models.Donor) error {

	querystmt, err := pr.conn.Prepare("SELECT username,password FROM donor WHERE username=$1 AND password=$2")

	if err != nil {
		return err
	}
	var username string
	var pass string

	err = querystmt.QueryRow(d.Username, d.Password).Scan(&username, &pass)

	if err == sql.ErrNoRows {
		return err
	}

	if err != nil {
		return err
	}
	
	return nil
}
func (pr *PsqlDonorRepository) SelectAllDonor() (donors []models.Donor, err error) {
	selRec, err := pr.conn.Query("SELECT id ,firstname, lastname, address,occupation,phone,email FROM donor")
	if err != nil {
		return donors, errors.New("something")
	}
	donor := models.Donor{}
	for selRec.Next() {
		err := selRec.Scan(&donor.DonorNo, &donor.FirstName, &donor.LastName, &donor.Address, &donor.Occupation, &donor.PhoneNumber, &donor.EmailAddress)
		if err != nil {
			return donors, errors.New("Couldnot")
		}
		donors = append(donors, donor)
	}
	return donors, nil
}
