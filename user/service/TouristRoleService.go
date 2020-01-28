package service

import (
	"github.com/gplus1/HouseRentalSystem//entity/Tourists"
	"github.com/gplus1/HouseRentalSystem//entity"
)


type RoleService struct {
	roleRepo Tourist.RoleRepository
}


func NewRoleService(RoleRepo Tourist.RoleRepository) Tourist.RoleService {
	return &RoleService{roleRepo: RoleRepo}
}


func (TrS *RoleService) Roles() ([]entity.Role, []error) {

	rls, erTrS := TrS.roleRepo.Roles()
	if len(erTrS) > 0 {
		return nil, erTrS
	}
	return rls, erTrS

}


func (TrS *RoleService) Role(id uint) (*entity.Role, []error) {
	rl, erTrS := TrS.roleRepo.Role(id)
	if len(erTrS) > 0 {
		return nil, erTrS
	}
	return rl, erTrS

}


func (TrS *RoleService) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	rl, erTrS := TrS.roleRepo.UpdateRole(role)
	if len(erTrS) > 0 {
		return nil, erTrS
	}
	return rl, erTrS

}


func (TrS *RoleService) DeleteRole(id uint) (*entity.Role, []error) {

	rl, erTrS := TrS.roleRepo.DeleteRole(id)
	if len(erTrS) > 0 {
		return nil, erTrS
	}
	return rl, erTrS
}


func (TrS *RoleService) StoreRole(role *entity.Role) (*entity.Role, []error) {

	rl, erTrS := TrS.roleRepo.StoreRole(role)
	if len(erTrS) > 0 {
		return nil, erTrS
	}
	return rl, erTrS
}
