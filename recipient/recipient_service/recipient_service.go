package recipient_service

import (
	"fmt"
	"github.com/FundStation/models"
	"github.com/FundStation/recipient"
)



type RecipientService struct {
	rRepo recipient.RecipientRepository
}


func NewRecipientService(reciRepo recipient.RecipientRepository) *RecipientService {
	return &RecipientService{rRepo: reciRepo}
}

func (rs *RecipientService) SignupRecipient(recipient *models.Recipient) (*models.Recipient,error) {

	rec,err := rs.rRepo.InsertRecipient(recipient)

	if err != nil {
		return rec,err
	}

	return rec,nil
}

func (rs *RecipientService) LoginRecipient(recipient models.Recipient) error {

	err := rs.rRepo.SelectRecipient(recipient)

	if err != nil {
		return err
	}

	return nil
}
func (rs *RecipientService) ViewAllRecipient() (recipients []models.Recipient, err error) {

	recipients, err = rs.rRepo.SelectAllRecipient()

	if err != nil {
		return recipients, err
	}

	return recipients, nil
}

func (rs *RecipientService) PhoneExists(phone string) bool {

	isPhone := rs.rRepo.PhoneExists(phone)
	fmt.Println(isPhone)
	return isPhone


}
func (rs *RecipientService) UsernameExists(username string) bool {
	isUser:=  rs.rRepo.UsernameExists(username)
	return isUser
}


func (rs *RecipientService) EmailExists(email string) bool {
	isEmail:=  rs.rRepo.EmailExists(email)
	return isEmail
}

func (rs *RecipientService) RecipientByUsername(username string) (*models.Recipient,error) {
	recp,err:= rs.rRepo.RecipientByUsername(username)
	if err != nil{
		return recp,err
	}
	return recp,nil
}

func (rs *RecipientService) RecipientById(id int) (*models.Recipient,error) {
	recipient,err:=rs.rRepo.RecipientById(id)
	if err != nil{
		return recipient,err
	}
	return recipient,nil
}

func (rs *RecipientService) UpdateRecipientById(recipient *models.Recipient) (error) {
	err:=rs.rRepo.UpdateRecipientById(recipient)
	if err != nil{
		return err
	}
	return nil
}
func (rs *RecipientService) DeleteRecipientById(recipient *models.Recipient) (error) {
	err:=rs.rRepo.DeleteRecipientById(recipient)
	if err != nil{
		return err
	}
	return nil
}

func (rs *RecipientService) SelectByUsername(username string) (*models.RecipientInfo,error) {
	recp,err:=rs.rRepo.SelectByUsername(username)
	if err != nil{
		return recp,err
	}
	return recp,nil
}



