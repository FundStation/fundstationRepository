package admin

import "github.com/FundStation2/models"

type AdminRepository interface {
	SelectAdmin(string) (*models.Admin,error)
}
