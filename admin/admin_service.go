package admin

import "github.com/FundStation/models"

type AdminService interface {
	LoginAdmin(string) (*models.Admin,error)
}
