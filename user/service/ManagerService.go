package service

import (
	"github.com/gplus1/HouseRentalSystem//Manager"
)

type ManagerService struct {
	ManagerRepo Manager.ManagerRepository
}

func NewManagerService(ManagerRepository Manager.ManagerRepository) Manager.ManagerService {
	return &ManagerService{ManagerRepo: ManagerRepository}
}

func (us *ManagerService) Managers() ([]Manager.Manager, []error) {
	usrs, errs := us.ManagerRepo.Managers()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

func (us *ManagerService) Manager(id uint) (*Manager.Manager, []error) {
	usr, errs := us.ManagerRepo.Manager(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (us *ManagerService) UpdateManager(Manager *Manager.Manager) (*Manager.Manager, []error) {
	usr, errs := us.ManagerRepo.UpdateManager(Manager)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (us *ManagerService) DeleteManager(id uint) (*Manager.Manager, []error) {
	usr, errs := us.ManagerRepo.DeleteManager(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (us *ManagerService) StoreManager(Manager *Manager.Manager) (*Manager.Manager, []error) {
	usr, errs := us.ManagerRepo.StoreManager(Manager)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
