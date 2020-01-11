package user

import "github.com/gplus1/HouseRentalSystem/Back_End/entity/users"

type TouristRepository interface {
	Tourists() ([]entity.Tourist, []error)
	Tourist(id uint) (*entity.Tourist, []error)
	UpdateTourist(Tourist *entity.Tourist) (*entity.Tourist, []error)
	DeleteTourist(id uint) (*entity.Tourist, []error)
	StoreTourist(Tourist *entity.Tourist) (*entity.Tourist, []error)
}

type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}
