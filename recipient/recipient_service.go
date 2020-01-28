package recipient

import "github.com/FundStation/models"

type RecipientService interface {
	SignupRecipient(recipient *models.Recipient)(*models.Recipient,error)
	LoginRecipient(recipient models.Recipient) error
	ViewAllRecipient() ([]models.Recipient, error)
	RecipientByUsername(username string)(*models.Recipient,error)
	RecipientById(id int) (*models.Recipient,error)
	SelectByUsername(string) ( *models.RecipientInfo, error)
	UpdateRecipientById(donor *models.Recipient) (error)
	DeleteRecipientById(donor *models.Recipient) (error)
	UsernameExists(username string) bool
	EmailExists(email string) bool
	PhoneExists(phone string) bool

}
