package recipient_repository

import (
	"database/sql"
	"errors"
	"github.com/FundStation/models"
	"github.com/jinzhu/gorm"
)


type PsqlRecipientRepository struct {
	conn *sql.DB
}


func NewPsqlRecipientRepository(Conn *sql.DB) *PsqlRecipientRepository {
	return &PsqlRecipientRepository{conn: Conn}
}

func (pr *PsqlRecipientRepository) InsertRecipient(r *models.Recipient)(*models.Recipient, error) {

	err := pr.conn.QueryRow("INSERT INTO recipient(firstname,lastname,address,occupation,username,password,phone,email,role_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8,$9) returning id", r.FirstName, r.LastName, r.Address, r.Occupation, r.Username, r.Password, r.PhoneNumber, r.EmailAddress,r.RoleID).Scan(&r.RecipientNo)
	if err != nil {
		return r,err
	}
	return r,nil
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

	//_, err := pr.conn.Exec("SELECT username,password FROM recipient WHERE username=$1, password=$2", r.Username, r.Password)
	return nil
}
func (pr *PsqlRecipientRepository) SelectAllRecipient() (recipients []models.Recipient, err error) {
	selRec, err := pr.conn.Query("SELECT id ,firstname, lastname, address,occupation,phone,email FROM recipient")
	if err != nil {
		return recipients, errors.New("something")
	}
	recp := models.Recipient{}
	for selRec.Next() {
		err := selRec.Scan(&recp.RecipientNo, &recp.FirstName, &recp.LastName, &recp.Address, &recp.Occupation, &recp.PhoneNumber, &recp.EmailAddress)
		if err != nil {
			return recipients, errors.New("Couldnot")
		}
		recipients = append(recipients, recp)
	}
	return recipients, nil
}
func (pr *PsqlRecipientRepository) PhoneExists(phone string) bool {



	err := pr.conn.QueryRow("SELECT * FROM recipient WHERE phone=$1",phone)

	if err != nil {
		return false
	}

	return true
}
func (pr *PsqlRecipientRepository) UsernameExists(username string) bool {
	err := pr.conn.QueryRow("SELECT * FROM recipient WHERE username=$1",username)

	if err != nil {
		return false
	}

	return true
}

// EmailExists check if a given email is found
func (pr *PsqlRecipientRepository) EmailExists(email string) bool {
	err := pr.conn.QueryRow("SELECT * FROM recipient WHERE email=$1", email)

	if err != nil {
		return false
	}

	return true
}

func (pr *PsqlRecipientRepository) RecipientByUsername(username string) (*models.Recipient,error) {
	recipient:=&models.Recipient{}
	querystmt, err := pr.conn.Prepare("SELECT * FROM recipient WHERE username=$1")

	if err != nil {
		return recipient,err
	}
	err = querystmt.QueryRow(username).Scan(&recipient.RecipientNo, &recipient.FirstName, &recipient.LastName, &recipient.Address, &recipient.Occupation,&recipient.Username,&recipient.Password, &recipient.PhoneNumber, &recipient.EmailAddress,&recipient.RoleID)

	return recipient,nil
}

























type RecipientGormRepo struct {
	conn *gorm.DB
}

// NewCommentGormRepo returns new object of CommentGormRepo
func NewRecipientGormRepo(db *gorm.DB) *RecipientGormRepo{
	return &RecipientGormRepo{conn: db}
}

// StoreComment stores a given customer comment in the database
func (recpRepo *RecipientGormRepo) InsertRecipientGorm(recp *models.Recipient) (*models.Recipient, []error) {
	recipient := recp
	errs := recpRepo.conn.Create(recipient).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return recipient, errs
}


// Comments returns all customer comments stored in the database
func (recpRepo *RecipientGormRepo) SelectAllRecipientGorm() ([]models.Recipient, []error) {
	recipient := []models.Recipient{}
	errs := recpRepo.conn.Find(&recipient).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return recipient, errs
}

// Comment retrieves a customer comment from the database by its id
func (recpRepo *RecipientGormRepo) SelectRecipientGorm(username string) (*models.Recipient, []error) {
	recipient := models.Recipient{}
	errs := recpRepo.conn.First(&recipient,"username=?", username).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &recipient, errs
}

