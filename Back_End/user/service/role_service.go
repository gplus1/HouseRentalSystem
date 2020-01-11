package service

import (
	import "github.com/gplus1/HouseRentalSystem/Back_End/entity/users"
	import "github.com/gplus1/HouseRentalSystem/Back_End/entity"
)


type RoleService struct {
	roleRepo user.RoleRepository
}


func NewRoleService(RoleRepo user.RoleRepository) user.RoleService {
	return &RoleService{roleRepo: RoleRepo}
}


func (rs *RoleService) Roles() ([]entity.Role, []error) {

	rls, errs := rs.roleRepo.Roles()
	if len(errs) > 0 {
		return nil, errs
	}
	return rls, errs

}


func (rs *RoleService) Role(id uint) (*entity.Role, []error) {
	rl, errs := rs.roleRepo.Role(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs

}


func (rs *RoleService) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	rl, errs := rs.roleRepo.UpdateRole(role)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs

}


func (rs *RoleService) DeleteRole(id uint) (*entity.Role, []error) {

	rl, errs := rs.roleRepo.DeleteRole(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}


func (rs *RoleService) StoreRole(role *entity.Role) (*entity.Role, []error) {

	rl, errs := rs.roleRepo.StoreRole(role)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}
