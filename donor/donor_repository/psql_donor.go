package donor_repository

import (
	"database/sql"
	"errors"
	"fmt"

	//"errors"

	"github.com/FundStation/models"
)


type PsqlDonorRepository struct {
	conn *sql.DB
}


func NewPsqlDonorRepository(Conn *sql.DB) *PsqlDonorRepository {
	return &PsqlDonorRepository{conn: Conn}
}

func (pr *PsqlDonorRepository) InsertDonor(d *models.Donor) (*models.Donor,error) {

	_, err := pr.conn.Exec("INSERT INTO donor(firstname,lastname,address,occupation,username,password,phone,email,role_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8,$9)", d.FirstName, d.LastName, d.Address, d.Occupation, d.Username, d.Password, d.PhoneNumber, d.EmailAddress,d.RoleID)
	if err != nil {
		return d,err
	}
	return d,nil
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

	//_, err := pr.conn.Exec("SELECT username,password FROM recipient WHERE username=$1, password=$2", r.Username, r.Password)
	return nil
}

func (pr *PsqlDonorRepository) SelectAllDonor() (donors []models.Donor, err error) {
	selRec, err := pr.conn.Query("SELECT * FROM donor")
	if err != nil {
		fmt.Println(err)
		return donors, errors.New("something")

	}

	donor := models.Donor{}
	for selRec.Next() {
		err := selRec.Scan(&donor.FirstName, &donor.LastName, &donor.Address, &donor.Occupation,&donor.Username,&donor.Password, &donor.PhoneNumber, &donor.EmailAddress,&donor.DonorNo,&donor.RoleID)
		if err != nil {
			fmt.Println("err",err)
			return donors, errors.New("Couldnot")
		}
		donors = append(donors, donor)

	}

	return donors, nil
}

func (pr *PsqlDonorRepository) PhoneExists(phone string) bool {
	var name string
	err := pr.conn.QueryRow("SELECT firstname FROM donor WHERE phone=$1",phone).Scan(&name)
	if err != nil {
		return false
	}

	return true
}
func (pr *PsqlDonorRepository) UsernameExists(username string) bool {
	var name string
	err := pr.conn.QueryRow("SELECT firstname FROM donor WHERE username=$1",username).Scan(&name)
	if err != nil {
		return false
	}

	return true
}

// EmailExists check if a given email is found
func (pr *PsqlDonorRepository) EmailExists(email string) bool {
	var name string
	err := pr.conn.QueryRow("SELECT firstname FROM donor WHERE email=$1", email).Scan(&name)
	if err != nil {
		return false
	}

	return true
}

func (pr *PsqlDonorRepository) DonorByUsername(username string) (*models.Donor,error) {
	donor:=&models.Donor{}
	querystmt, err := pr.conn.Prepare("SELECT * FROM donor WHERE username=$1")

	if err != nil {
		return donor,err
	}
	err = querystmt.QueryRow(username).Scan(&donor.FirstName, &donor.LastName, &donor.Address, &donor.Occupation,&donor.Username,&donor.Password, &donor.PhoneNumber, &donor.EmailAddress,&donor.DonorNo,&donor.RoleID)
	if err != nil{
		return donor,err
	}

	return donor,nil
}
func (pr *PsqlDonorRepository) DonorById(id int) (*models.Donor,error) {
	donor:=&models.Donor{}
	querystmt, err := pr.conn.Prepare("SELECT * FROM donor WHERE id=$1")

	if err != nil {
		fmt.Println(err)
		return donor,err
	}
	err = querystmt.QueryRow(id).Scan(&donor.FirstName, &donor.LastName, &donor.Address, &donor.Occupation,&donor.Username,&donor.Password, &donor.PhoneNumber, &donor.EmailAddress,&donor.DonorNo,&donor.RoleID)
	if err != nil{
		fmt.Println("error",err)
	}
	fmt.Println("donr",donor)

	return donor,nil
}
func (pr *PsqlDonorRepository) UpdateDonorById(donor *models.Donor) (error) {

	_, err := pr.conn.Exec("UPDATE donor set lastname = $2 where id=$1",donor.DonorNo,donor.LastName)

	if err != nil {
		fmt.Println(err)
		return err
	}


	return nil
}
func (pr *PsqlDonorRepository) DeleteDonorById(donor *models.Donor) (error) {

	_, err := pr.conn.Exec("DELETE FROM donor where id=$1",donor.DonorNo)

	if err != nil {
		fmt.Println(err)
		return err
	}


	return nil
}
