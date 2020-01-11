package donor

import "github.com/FundStation/models"

type DonorService interface {
	SignupDonor(donor *models.Donor) error
	LoginDonor(donor models.Donor) error
	ViewAllDonor() ([]models.Donor, error)
}
