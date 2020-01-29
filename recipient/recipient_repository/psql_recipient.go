package recipient_repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/FundStation/models"
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

	var name string
	err := pr.conn.QueryRow("SELECT firstname FROM recipient WHERE phone=$1",phone).Scan(&name)

	if err != nil  {
		return false
	}
	return true
}

func (pr *PsqlRecipientRepository) UsernameExists(username string) bool {
	var name string
	err:= pr.conn.QueryRow("SELECT firstname FROM recipient WHERE username=$1",username).Scan(&name)
	if err != nil {
		return false
	}

	return true
}

// EmailExists check if a given email is found
func (pr *PsqlRecipientRepository) EmailExists(email string) bool {
	var name string
	err := pr.conn.QueryRow("SELECT firstname FROM recipient WHERE email=$1",email).Scan(&name)

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

func (pr *PsqlRecipientRepository) RecipientById(id int) (*models.Recipient,error) {
	recipient:=&models.Recipient{}
	querystmt, err := pr.conn.Prepare("SELECT * FROM recipient WHERE id=$1")

	if err != nil {
		fmt.Println(err)
		return recipient,err
	}
	err = querystmt.QueryRow(id).Scan(&recipient.RecipientNo,&recipient.FirstName, &recipient.LastName, &recipient.Address, &recipient.Occupation,&recipient.Username,&recipient.Password, &recipient.PhoneNumber, &recipient.EmailAddress,&recipient.RoleID)
	if err != nil{
		fmt.Println("error",err)
	}
	fmt.Println("donr",recipient)

	return recipient,nil
}
func (pr *PsqlRecipientRepository) UpdateRecipientById(recipient *models.Recipient) (error) {

	_, err := pr.conn.Exec("UPDATE recipient set lastname = $2 where id=$1",recipient.RecipientNo,recipient.LastName)

	if err != nil {
		fmt.Println(err)
		return err
	}


	return nil
}
func (pr *PsqlRecipientRepository) DeleteRecipientById(recipient *models.Recipient) (error) {

	_, err := pr.conn.Exec("DELETE FROM recipient where id=$1",recipient.RecipientNo)

	if err != nil {
		fmt.Println(err)
		return err
	}


	return nil
}
func (pr *PsqlRecipientRepository) SelectByUsername(username string) ( *models.RecipientInfo, error){
	recipientInfo := &models.RecipientInfo{}
	recp := &models.Recipient{}
	err := pr.conn.QueryRow("SELECT firstname,lastname,image,description  FROM recipient INNER JOIN recipientinfo ON recipientinfo.recipient_id = recipient.id  WHERE recipient.username =$1",username).Scan(&recp.FirstName,&recp.LastName,&recipientInfo.Image,&recipientInfo.Description)
	recipientInfo.Recipient=recp
	if err != nil {
		fmt.Println("Erro",err)

	}

	return recipientInfo,nil
}

