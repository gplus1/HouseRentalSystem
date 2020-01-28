package Tourist

import (
	"github.com/gplus1/HouseRentalSystem//entity/Tourists"
 "github.com/gplus1/HouseRentalSystem//entity"
)

type ManagerRepository interface {
	Managers() ([]Manager.Manager, []error)
	Manager(id uint) (*Manager.Manager, []error)
	UpdateManager(Manager *Manager.Manager) (*Manager.Manager, []error)
	DeleteManager(id uint) (*Manager.Manager, []error)
	StoreManager(Manager *Manager.Manager) (*Manager.Manager, []error)
}


type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}