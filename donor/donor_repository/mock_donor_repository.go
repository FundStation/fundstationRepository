package donor_repository

import (
	"database/sql"
	"errors"
	"github.com/FundStation2/donor"
	"github.com/FundStation2/models"
)


type MockDonorRepo struct {
	conn *sql.DB
}




func NewMockDonorRepo(db *sql.DB) donor.DonorRepository {
	return &MockDonorRepo{conn: db}
}


func (mDonRepo *MockDonorRepo)SelectAllDonor() ([]models.Donor, error) {
	donors := []models.Donor{models.DonorMock}
	return donors, nil
}


func (mDonRepo *MockDonorRepo) SelectDonor(donor models.Donor) (error) {
	donor = models.DonorMock
	if donor.Username == "mockyuser" && donor.Password=="password" {
		return nil
	}
	return errors.New("Not found")
}

func (mDonRepo *MockDonorRepo) InsertDonor(donor *models.Donor) (*models.Donor,error) {
	don := donor
	return don,nil
}

func (mDonRepo *MockDonorRepo) PhoneExists(phone string) bool {

	if phone != models.DonorMock.PhoneNumber{
		return false
	}

	return true
}
func (mDonRepo *MockDonorRepo) UsernameExists(username string) bool {
	if username != models.DonorMock.Username{
		return false
	}

	return true
}


func (mDonRepo *MockDonorRepo) EmailExists(email string) bool {


	if email != models.DonorMock.EmailAddress{
		return false
	}

	return true
}

func (mDonRepo *MockDonorRepo) DonorByUsername(username string) (*models.Donor,error) {

	if username == "mockyuser" {

		return &models.DonorMock,nil
	}
	return nil,errors.New("Not found")
}
func (mDonRepo *MockDonorRepo) DonorById(id int) (*models.Donor,error) {

	if id == 1 {

		return &models.DonorMock,nil
	}
	return nil,errors.New("Not found")
}

func (mDonRepo *MockDonorRepo) UpdateDonorById(donor *models.Donor) (error) {

	return nil
}

func (mDonRepo *MockDonorRepo) DeleteDonorById(donor *models.Donor) (error) {

return nil
}
