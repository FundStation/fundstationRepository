package recipientInfo_service

import (
	"fmt"
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
func (ris *RecipientInfoService) AccountExistsInfo(account string) bool {
	isAccountInfo:= ris.rInfoRepo.AccountExistsInfo(account)
	return isAccountInfo
}

func (ris *RecipientInfoService) SelectApproved() (dInfo []models.DonationInfo,err error){
	dInfo,err=ris.rInfoRepo.SelectApproved()
	if(err != nil){
		return dInfo,err

	}
	return dInfo,nil

}
func (ris *RecipientInfoService) SelectIndividualById(id int) (dInfo models.DonationInfo,err error){
	dInfo,err=ris.rInfoRepo.SelectIndividualById(id)
	if(err != nil){
		return dInfo,err

	}
	return dInfo,nil

}

func (ris  *RecipientInfoService) DeleteRecipientInfoById(id int) (error) {

	err := ris.rInfoRepo.DeleteRecipientInfoById(id)

	if err != nil {
		fmt.Println(err)
		return err
	}


	return nil
}

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
func (ris *RecipientInfoService) AccountExistsInfo(account string) bool {
	isAccountInfo:= ris.rInfoRepo.AccountExistsInfo(account)
	return isAccountInfo
}
func (ris *RecipientInfoService) SelectApproved() (rinfo []models.RecipientInfo,err error){
	rinfo,err=ris.rInfoRepo.SelectApproved()
	if(err != nil){
		return rinfo,err

	}
	return rinfo,nil

}
