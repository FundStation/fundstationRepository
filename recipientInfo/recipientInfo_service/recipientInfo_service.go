package recipientInfo_service

import (
	"github.com/FundStation/models"
	"github.com/FundStation/recipientInfo"
)

type RecipientInfoService struct {
	rInfoRepo recipientInfo.RecipientInfoRepository
}

func NewRecipientInfoService(reciInfoRepo recipientInfo.RecipientInfoRepository) *RecipientInfoService {
	return &RecipientInfoService{rInfoRepo: reciInfoRepo}
}

func (ris *RecipientInfoService) CreateRecipientInfo(recipientInfo models.RecipientInfo) error {

	err := ris.rInfoRepo.InsertRecipientInfo(recipientInfo)

	if err != nil {
		return err
	}

	return nil
}

func (ris *RecipientInfoService) ViewSpecificRecipientInfo(id int) (recipientInfo models.RecipientInfo, err error) {

	recipientInfo, err = ris.rInfoRepo.SelectRecipientInfo(id)

	if err != nil {
		return recipientInfo, err
	}

	return recipientInfo, nil
}
func (ris *RecipientInfoService) ApproveRecipientInfo(id int) (err error) {

	err = ris.rInfoRepo.UpdateRecipientInfo(id)

	if err != nil {
		return err
	}

	return nil
}
