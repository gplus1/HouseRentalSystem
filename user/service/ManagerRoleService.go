package service

import (
	"github.com/gplus1/HouseRentalSystem//entity"
	"github.com/gplus1/HouseRentalSystem//entity/Tourists/Manager"
)

type RoleService struct {
	roleRepo Manager.RoleRepository
}

func NewRoleService(RoleRepo Manager.RoleRepository) Manager.RoleService {
	return &RoleService{roleRepo: RoleRepo}
}

/MrS stands for Manager Role Service
func (MrS *RoleService) Roles() ([]entity.Role, []error) {

	rls, erMrS := MrS.roleRepo.Roles()
	if len(erMrS) > 0 {
		return nil, erMrS
	}
	return rls, erMrS

}

func (MrS *RoleService) Role(id uint) (*entity.Role, []error) {
	rl, erMrS := MrS.roleRepo.Role(id)
	if len(erMrS) > 0 {
		return nil, erMrS
	}
	return rl, erMrS

}

func (MrS *RoleService) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	rl, erMrS := MrS.roleRepo.UpdateRole(role)
	if len(erMrS) > 0 {
		return nil, erMrS
	}
	return rl, erMrS

}

func (MrS *RoleService) DeleteRole(id uint) (*entity.Role, []error) {

	rl, erMrS := MrS.roleRepo.DeleteRole(id)
	if len(erMrS) > 0 {
		return nil, erMrS
	}
	return rl, erMrS
}

func (MrS *RoleService) StoreRole(role *entity.Role) (*entity.Role, []error) {

	rl, erMrS := MrS.roleRepo.StoreRole(role)
	if len(erMrS) > 0 {
		return nil, erMrS
	}
	return rl, erMrS
}
