package recipient

import "github.com/FundStation/models"

type RecipientService interface {
	SignupRecipient(recipient *models.Recipient) error
	LoginRecipient(recipient models.Recipient) error
	ViewAllRecipient() ([]models.Recipient, error)
}
