package service

import (
	"github.com/gplus1/HouseRentalSystem//entity/Tourists"
	"github.com/gplus1/HouseRentalSystem//entity"
)

type LandLordService interface {
	LandLords() ([]entity.LandLord, []error)
	LandLord(id uint) (*entity.LandLord, []error)
	UpdateLandLord(LandLord *entity.LandLord) (*entity.LandLord, []error)
	DeleteLandLord(id uint) (*entity.LandLord, []error)
	StoreLandLord(LandLord *entity.LandLord) (*entity.LandLord, []error)
}
type ManagerService interface {
	Managers() ([]entity.Manager, []error)
	Manager(id uint) (*entity.Manager, []error)
	UpdateManager(Manager *entity.Manager) (*entity.Manager, []error)
	DeleteManager(id uint) (*entity.Manager, []error)
	StoreManager(Manager *entity.Manager) (*entity.Manager, []error)
}
type LandLordService interface {
	LandLords() ([]entity.LandLord, []error)
	LandLord(id uint) (*entity.LandLord, []error)
	UpdateLandLord(LandLord *entity.LandLord) (*entity.LandLord, []error)
	DeleteLandLord(id uint) (*entity.LandLord, []error)
	StoreLandLord(LandLord *entity.LandLord) (*entity.LandLord, []error)
}
type TouristService interface {
	Tourists() ([]entity.Tourist, []error)
	Tourist(id uint) (*entity.Tourist, []error)
	UpdateTourist(Tourist *entity.Tourist) (*entity.Tourist, []error)
	DeleteTourist(id uint) (*entity.Tourist, []error)
	StoreTourist(Tourist *entity.Tourist) (*entity.Tourist, []error)
}

type RoleService interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}
type SessionService interface {
	Session(sessionID string) (*entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionID string) (*entity.Session, []error)
}