package user

import "github.com/gplus1/HouseRentalSystem/Back_End/entity/users"
import "github.com/betsegawlemma/restaurant/entity/users"

type AdminRepository interface {
	Admins() ([]entity.Admin, []error)
	Admin(id uint) (*entity.Admin, []error)
	UpdateAdmin(Admin *entity.Admin) (*entity.Admin, []error)
	DeleteAdmin(id uint) (*entity.Admin, []error)
	StoreAdmin(Admin *entity.Admin) (*entity.Admin, []error)
}


type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}
