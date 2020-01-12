package admin_service

import (
	"github.com/FundStation/admin"
)

// CategoryService implements menu.CategoryService interface
type AdminService struct {
	aRepo admin.AdminRepository
}

// NewCategoryService will create new CategoryService object
func NewAdminService(adRepo admin.AdminRepository) *AdminService {
	return &AdminService{aRepo: adRepo}
}
