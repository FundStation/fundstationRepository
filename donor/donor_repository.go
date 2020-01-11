package donor

import "github.com/FundStation/models"

type DonorRepository interface {
	InsertDonor(donor *models.Donor) error
	SelectDonor(donor models.Donor) error
	SelectAllDonor() ([]models.Donor, error)
}
