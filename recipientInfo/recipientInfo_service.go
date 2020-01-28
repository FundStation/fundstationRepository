package recipientInfo

import "github.com/FundStation/models"

type RecipientInfoService interface {
	CreateRecipientInfo(recipientInfo models.RecipientInfo) error
	ViewSpecificRecipientInfo(int) (models.RecipientInfo, error)
	ApproveRecipientInfo(int) error
	AccountExistsInfo(account string) bool
	SelectApproved() ( []models.DonationInfo, error)
	SelectIndividualById(id int) ( models.DonationInfo, error)
	DeleteRecipientInfoById(int) (error)
}
