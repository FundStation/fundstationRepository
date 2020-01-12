package recipientInfo

import "github.com/FundStation/models"

type RecipientInfoRepository interface {
	InsertRecipientInfo(recipientInfo models.RecipientInfo) error
	SelectRecipientInfo(int) (models.RecipientInfo, error)
	UpdateRecipientInfo(int) error
	AccountExistsInfo(account string) bool
	SelectApproved() (rinfo []models.RecipientInfo,err error)
}
