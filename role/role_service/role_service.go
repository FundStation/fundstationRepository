package role_service

import (
	"github.com/FundStation/models"
	"github.com/FundStation/role"
)


type RoleService struct {
	roleRepo role.RoleRepository
}


func NewRoleService(RoleRepo role.RoleRepository) *RoleService {
	return &RoleService{roleRepo: RoleRepo}
}


func (rs *RoleService) Roles() ([]models.Role, error) {

	rls, err := rs.roleRepo.Roles()
	if err != nil {
		return nil, err
	}
	return rls, err

}


func (rs *RoleService) RoleByName(name string) (*models.Role, error) {
	role, err := rs.roleRepo.RoleByName(name)
	if err != nil {
		return nil, err
	}
	return role, err
}


func (rs *RoleService) Role(id uint) (*models.Role, error) {
	rl, err := rs.roleRepo.Role(id)
	if err != nil {
		return nil, err
	}
	return rl, err

}


func (rs *RoleService) UpdateRole(role *models.Role) ( error) {
	err := rs.roleRepo.UpdateRole(role)
	if err != nil {
		return err
	}
	return  nil

}


func (rs *RoleService) DeleteRole(id uint) ( error) {

	 err := rs.roleRepo.DeleteRole(id)
	if err != nil {
		return  err
	}
	return nil
}


func (rs *RoleService) StoreRole(role *models.Role) ( error) {

	err := rs.roleRepo.StoreRole(role)
	if err != nil {
		return  err
	}
	return nil
}

func (rs *RoleService) DonorRoles(donor *models.Donor) (models.Role, error) {
	donorRoles, err := rs.roleRepo.DonorRoles(donor)
	if err != nil {
		return donorRoles, err
	}
	return donorRoles, err
}

func (rs *RoleService) RecipientRoles(recipient *models.Recipient) (models.Role, error) {
	recpRoles, err := rs.roleRepo.RecipientRoles(recipient)
	if err != nil {
		return recpRoles, err
	}
	return recpRoles, err
}
