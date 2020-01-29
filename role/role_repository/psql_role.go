package role_repository

import (
	"database/sql"
	"errors"
	"github.com/FundStation/models"
)


type RoleRepository struct {
	conn *sql.DB
}


func NewRoleRepository(Conn *sql.DB) *RoleRepository {
	return &RoleRepository{conn: Conn}
}


func (rr *RoleRepository) StoreRole(role *models.Role) error {

	query := "INSERT INTO roles (name) values ($1)"
	_, err := rr.conn.Exec(query, role.Name)

	if err != nil {
		return errors.New("INSERT has failed")
	}

	return nil
}

func (rr *RoleRepository) RoleByName(name string) (*models.Role,error) {
	role :=&models.Role{}
	querystmt, err := rr.conn.Prepare("SELECT * FROM roles WHERE name=$1")

	if err != nil {
		return role, err
	}

	err = querystmt.QueryRow(name).Scan(&role.ID, &role.Name)
	if err != nil {
		return role, err
	}

	return role, nil
}
func (rr *RoleRepository) DonorRoles(donor *models.Donor) (models.Role, error) {
	donorRole := models.Role{}

	querystmt, err := rr.conn.Prepare("SELECT * FROM roles WHERE id=$1")

	if err != nil {
		return donorRole, err
	}

	err = querystmt.QueryRow(donor.RoleID).Scan(&donorRole.Name, &donorRole.ID)
	if err != nil {
		return donorRole, err
	}

	return donorRole, nil
}
func (rr *RoleRepository) RecipientRoles(recipient *models.Recipient) (models.Role, error) {
	recpRole := models.Role{}

	querystmt, err := rr.conn.Prepare("SELECT * FROM roles WHERE id=$1")

	if err != nil {
		return recpRole, err
	}

	err = querystmt.QueryRow(recipient.RoleID).Scan(&recpRole.Name, &recpRole.ID)
	if err != nil {
		return recpRole, err
	}

	return recpRole, nil
}
func (rr *RoleRepository)AdminRoles(admin *models.Admin) (models.Role, error) {
	recpRole := models.Role{}

	querystmt, err := rr.conn.Prepare("SELECT * FROM roles WHERE id=$1")

	if err != nil {
		return recpRole, err
	}

	err = querystmt.QueryRow(3).Scan(&recpRole.ID, &recpRole.Name)
	if err != nil {
		return recpRole, err
	}

	return recpRole, nil
}
func (rr *RoleRepository) Roles() ([]models.Role, error) {
	return []models.Role{}, nil
}

func (rr *RoleRepository) Role(id uint) (*models.Role, error) {
	return &models.Role{}, nil
}

func (rr *RoleRepository) UpdateRole(role *models.Role) error { return nil }

func (rr *RoleRepository) DeleteRole(id uint) error { return nil }
