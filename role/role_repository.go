package role

import "github.com/FundStation2/models"

type RoleRepository interface {
	Roles() ([]models.Role, error)
	Role(id uint) (*models.Role, error)
	RoleByName(name string) (*models.Role, error)
	UpdateRole(role *models.Role) ( error)
	DeleteRole(id uint) (error)
	StoreRole(role *models.Role) (error)
	DonorRoles(donor *models.Donor) (models.Role, error)
	RecipientRoles(donor *models.Recipient) (models.Role, error)
	AdminRoles(donor *models.Admin) (models.Role, error)
}

type SessionRepository interface {
	Session(sessionID string) (*models.Session, error)
	StoreSession(session *models.Session) (*models.Session, error)
	DeleteSession(sessionID string) (*models.Session, error)
}
