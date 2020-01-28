package Tourist

import "github.com/gplus1/HouseRentalSystem//entity/Tourists"
import "github.com/gplus1/HouseRentalSystem//entity"

type LandLordRepository interface {
	LandLords() ([]LandLord.LandLord, []error)
	LandLord(id uint) (*LandLord.LandLord, []error)
	UpdateLandLord(LandLord *LandLord.LandLord) (*LandLord.LandLord, []error)
	DeleteLandLord(id uint) (*LandLord.LandLord, []error)
	StoreLandLord(LandLord *LandLord.LandLord) (*LandLord.LandLord, []error)
}


type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}
