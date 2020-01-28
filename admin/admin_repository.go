package admin

import "github.com/FundStation/models"

type AdminRepository interface {
	SelectAdmin(string) (*models.Admin,error)
}
