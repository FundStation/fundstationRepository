package recipient

import "github.com/FundStation/models"

type RecipientRepository interface {
	InsertRecipient(recipient *models.Recipient) error
	SelectRecipient(recipient models.Recipient) error
	SelectAllRecipient() ([]models.Recipient, error)
}
