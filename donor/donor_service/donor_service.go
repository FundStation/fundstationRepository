package donor_service

import (
	"github.com/FundStation/donor"
	"github.com/FundStation/models"
)

type DonorService struct {
	dRepo donor.DonorRepository
}

func NewDonorService(donorRepo donor.DonorRepository) *DonorService {
	return &DonorService{dRepo: donorRepo}
}

func (ds *DonorService) SignupDonor(donor *models.Donor) error {

	err := ds.dRepo.InsertDonor(donor)

	if err != nil {
		return err
	}

	return nil
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
