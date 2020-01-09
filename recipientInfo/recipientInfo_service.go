package recipientInfo

import "github.com/FundStation/models"

// CategoryRepository specifies menu category related database operations
type RecipientInfoService interface {
	CreateRecipientInfo(recipientInfo models.RecipientInfo) error
	ViewSpecificRecipientInfo(int) (models.RecipientInfo, error)
	ApproveRecipientInfo(int) error
}
