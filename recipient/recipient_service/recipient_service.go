package recipient_service

import (
	"github.com/FundStation/models"
	"github.com/FundStation/recipient"
)

type RecipientService struct {
	rRepo recipient.RecipientRepository
}

func NewRecipientService(reciRepo recipient.RecipientRepository) *RecipientService {
	return &RecipientService{rRepo: reciRepo}
}

func (rs *RecipientService) SignupRecipient(recipient *models.Recipient) error {

	err := rs.rRepo.InsertRecipient(recipient)

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

func (rs *RecipientService) LoginRecipient(recipient models.Recipient) error {

	err := rs.rRepo.SelectRecipient(recipient)

	if err != nil {
		return err
	}

	return nil
}
