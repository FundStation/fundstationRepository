package donor

import "github.com/FundStation/models"


type DonorRepository interface {
	InsertDonor(donor *models.Donor) (*models.Donor,error)
	SelectDonor(donor models.Donor) error
	SelectAllDonor() ([]models.Donor, error)
	DonorByUsername(username string)(*models.Donor,error)
	DonorById(id int) (*models.Donor,error)
	UpdateDonorById(donor *models.Donor) (error)
	DeleteDonorById(donor *models.Donor) (error)
	UsernameExists(username string) bool
	EmailExists(email string) bool
	PhoneExists(phone string) bool
}
