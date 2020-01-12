package recipient

import (
	"github.com/FundStation/models"
)

type RecipientRepository interface {
	InsertRecipient(recipient *models.Recipient)(*models.Recipient, error)
	SelectRecipient(recipient models.Recipient) error
	SelectAllRecipient() ([]models.Recipient, error)
	RecipientByUsername(username string)(*models.Recipient,error)
	UsernameExists(username string) bool
	EmailExists(email string) bool
	PhoneExists(phone string) bool
}
