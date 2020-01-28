package admin_service

import (
	"github.com/FundStation/admin"
	"github.com/FundStation/models"
)


type AdminService struct {
	aRepo admin.AdminRepository
}


func NewAdminService(adRepo admin.AdminRepository) *AdminService {
	return &AdminService{aRepo: adRepo}
}
func (as *AdminService) LoginAdmin(user string) (*models.Admin,error) {

	adm,err := as.aRepo.SelectAdmin(user)

	if err != nil {
		return adm, err
	}

	return adm,nil
}
