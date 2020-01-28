package donor

import "github.com/FundStation/models"

type DonorService interface {
	SignupDonor(donor *models.Donor) (*models.Donor,error)
	LoginDonor(donor models.Donor) error
	ViewAllDonor() ([]models.Donor, error)
	DonorByUsername(username string)(*models.Donor,error)
	DonorById(id int) (*models.Donor,error)
	UpdateDonorById(donor *models.Donor) (error)
	DeleteDonorById(donor *models.Donor) (error)
	UsernameExists(username string) bool
	EmailExists(email string) bool
	PhoneExists(phone string) bool
}
