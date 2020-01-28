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

func (ArS *RoleService) Roles() ([]entity.Role, []error) {

	rls, erArS := ArS.roleRepo.Roles()
	if len(erArS) > 0 {
		return nil, erArS
	}
	return rls, erArS

}

func (ArS *RoleService) Role(id uint) (*entity.Role, []error) {
	rl, erArS := ArS.roleRepo.Role(id)
	if len(erArS) > 0 {
		return nil, erArS
	}
	return rl, erArS

}

func (ArS *RoleService) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	rl, erArS := ArS.roleRepo.UpdateRole(role)
	if len(erArS) > 0 {
		return nil, erArS
	}
	return rl, erArS

}

func (ArS *RoleService) DeleteRole(id uint) (*entity.Role, []error) {

	rl, erArS := ArS.roleRepo.DeleteRole(id)
	if len(erArS) > 0 {
		return nil, erArS
	}
	return rl, erArS
}

func (ArS *RoleService) StoreRole(role *entity.Role) (*entity.Role, []error) {

	rl, erArS := ArS.roleRepo.StoreRole(role)
	if len(erArS) > 0 {
		return nil, erArS
	}
	return rl, erArS
}
