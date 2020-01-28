package donor_service

import (
	"fmt"
	"github.com/FundStation/donor"
	"github.com/FundStation/models"
)


type DonorService struct {
	dRepo donor.DonorRepository
}


func NewDonorService(donorRepo donor.DonorRepository) *DonorService {
	return &DonorService{dRepo: donorRepo}
}

func (ds *DonorService) SignupDonor(donor *models.Donor) (*models.Donor,error) {

	don,err := ds.dRepo.InsertDonor(donor)

	if err != nil {
		return don,err
	}

	return don,nil
}
func (ds *DonorService) LoginDonor(donor models.Donor) error {

	err := ds.dRepo.SelectDonor(donor)

	if err != nil {
		return err
	}

	return nil
}

func (ds *DonorService) ViewAllDonor() (donors []models.Donor, err error) {

	donors, err = ds.dRepo.SelectAllDonor()

	if err != nil {
		return donors, err
	}

	return donors, nil
}

func (ds *DonorService) PhoneExists(phone string) bool {

	isPhone := ds.dRepo.PhoneExists(phone)
	fmt.Println(isPhone)
	return isPhone


}
func (ds *DonorService) UsernameExists(username string) bool {
	isUser:= ds.dRepo.UsernameExists(username)
	return isUser
}


func (ds *DonorService) EmailExists(email string) bool {
	isEmail:= ds.dRepo.EmailExists(email)
	return isEmail
}

func (ds *DonorService) DonorByUsername(username string) (*models.Donor,error) {
	donor,err:=ds.dRepo.DonorByUsername(username)
	if err != nil{
		return donor,err
	}
	return donor,nil
}
func (ds *DonorService) DonorById(id int) (*models.Donor,error) {
	donor,err:=ds.dRepo.DonorById(id)
	if err != nil{
		return donor,err
	}
	return donor,nil
}

func (ds *DonorService) UpdateDonorById(donor *models.Donor) (error) {
	err:=ds.dRepo.UpdateDonorById(donor)
	if err != nil{
		return err
	}
	return nil
}
func (ds *DonorService) DeleteDonorById(donor *models.Donor) (error) {
	err:=ds.dRepo.DeleteDonorById(donor)
	if err != nil{
		return err
	}
	return nil
}


