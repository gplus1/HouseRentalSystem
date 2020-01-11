package user

import "github.com/gplus1/HouseRentalSystem/Back_End/entity/users"

type LandLordRepository interface {
	LandLords() ([]entity.LandLord, []error)
	LandLord(id uint) (*entity.LandLord, []error)
	UpdateLandLord(LandLord *entity.LandLord) (*entity.LandLord, []error)
	DeleteLandLord(id uint) (*entity.LandLord, []error)
	StoreLandLord(LandLord *entity.LandLord) (*entity.LandLord, []error)
}

type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}
