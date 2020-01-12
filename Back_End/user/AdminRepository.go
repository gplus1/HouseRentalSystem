package user

import "github.com/gplus1/HouseRentalSystem/Back_End/entity/users"
import "github.com/gplus1/HouseRentalSystem/Back_End/entity"

type AdminRepository interface {
	Admins() ([]Admin.Admin, []error)
	Admin(id uint) (*Admin.Admin, []error)
	UpdateAdmin(Admin *Admin.Admin) (*Admin.Admin, []error)
	DeleteAdmin(id uint) (*Admin.Admin, []error)
	StoreAdmin(Admin *Admin.Admin) (*Admin.Admin, []error)
}


type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}
