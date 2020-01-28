package service

import (
	"github.com/gplus1/HouseRentalSystem//entity"
	"github.com/gplus1/HouseRentalSystem//entity/Tourists/LandLord"
)

type RoleService struct {
	roleRepo LandLord.RoleRepository
}

func NewRoleService(RoleRepo LandLord.RoleRepository) LandLord.RoleService {
	return &RoleService{roleRepo: RoleRepo}
}

func (LrS *RoleService) Roles() ([]entity.Role, []error) {

	rls, erLrS := LrS.roleRepo.Roles()
	if len(erLrS) > 0 {
		return nil, erLrS
	}
	return rls, erLrS

}

func (LrS *RoleService) Role(id uint) (*entity.Role, []error) {
	rl, erLrS := LrS.roleRepo.Role(id)
	if len(erLrS) > 0 {
		return nil, erLrS
	}
	return rl, erLrS

}

func (LrS *RoleService) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	rl, erLrS := LrS.roleRepo.UpdateRole(role)
	if len(erLrS) > 0 {
		return nil, erLrS
	}
	return rl, erLrS

}

func (LrS *RoleService) DeleteRole(id uint) (*entity.Role, []error) {

	rl, erLrS := LrS.roleRepo.DeleteRole(id)
	if len(erLrS) > 0 {
		return nil, erLrS
	}
	return rl, erLrS
}

func (LrS *RoleService) StoreRole(role *entity.Role) (*entity.Role, []error) {

	rl, erLrS := LrS.roleRepo.StoreRole(role)
	if len(erLrS) > 0 {
		return nil, erLrS
	}
	return rl, erLrS
}
