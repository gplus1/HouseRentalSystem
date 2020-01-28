package repository

import (
	"github.com/gplus1/HouseRentalSystem//Manager"
	"github.com/gplus1/HouseRentalSystem//Tourist"
	"github.com/jinzhu/gorm"
)

type ManagerGormRepo struct {
	conn *gorm.DB
}

func NewManagerGormRepo(db *gorm.DB) Manager.ManagerRepository {
	return &ManagerGormRepo{conn: db}
}

func (ManagerRepo *ManagerGormRepo) Managers() ([]entity.Manager, []error) {
	Managers := []entity.Manager{}
	errs := ManagerRepo.conn.Find(&Managers).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return Managers, errs
}

func (ManagerRepo *ManagerGormRepo) Manager(id uint) (*entity.Manager, []error) {
	Manager := entity.Manager{}
	errs := ManagerRepo.conn.First(&Manager, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &Manager, errs
}

func (ManagerRepo *ManagerGormRepo) UpdateManager(Manager *entity.Manager) (*entity.Manager, []error) {
	usr := Manager
	errs := ManagerRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (ManagerRepo *ManagerGormRepo) DeleteManager(id uint) (*entity.Manager, []error) {
	usr, errs := ManagerRepo.Manager(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = ManagerRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (ManagerRepo *ManagerGormRepo) StoreManager(Manager *entity.Manager) (*entity.Manager, []error) {
	usr := Manager
	errs := ManagerRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
