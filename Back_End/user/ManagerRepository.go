package user

import "github.com/gplus1/HouseRentalSystem/Back_End/entity/users"
import "github.com/betsegawlemma/restaurant/entity/users"

type ManagerRepository interface {
	Managers() ([]entity.Manager, []error)
	Manager(id uint) (*entity.Manager, []error)
	UpdateManager(Manager *entity.Manager) (*entity.Manager, []error)
	DeleteManager(id uint) (*entity.Manager, []error)
	StoreManager(Manager *entity.Manager) (*entity.Manager, []error)
}


type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}