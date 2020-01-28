package service

import (
	"github.com/gplus1/HouseRentalSystem//entity"
	"github.com/gplus1/HouseRentalSystem//entity/menu"
)

type H_status_service struct {
	HouseStatusRepo menu.HouseStatusRepository
}

func NewH_status_service(ingRepo menu.HouseStatusRepository) menu.H_status_service {
	return &H_status_service{HouseStatusRepo: ingRepo}
}

func (HsS *H_status_service) HouseStatuss() ([]entity.HouseStatus, []error) {
	ings, errs := HsS.HouseStatusRepo.HouseStatuss()
	if len(errs) > 0 {
		return nil, errs
	}
	return ings, errs
}

func (HsS *H_status_service) HouseStatus(id uint) (*entity.HouseStatus, []error) {
	ing, errs := HsS.HouseStatusRepo.HouseStatus(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return ing, errs
}

func (HsS *H_status_service) UpdateHouseStatus(HouseStatus *entity.HouseStatus) (*entity.HouseStatus, []error) {
	ing, errs := HsS.HouseStatusRepo.UpdateHouseStatus(HouseStatus)
	if len(errs) > 0 {
		return nil, errs
	}
	return ing, errs
}

func (HsS *H_status_service) DeleteHouseStatus(id uint) (*entity.HouseStatus, []error) {
	ing, errs := HsS.HouseStatusRepo.DeleteHouseStatus(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return ing, errs
}

func (HsS *H_status_service) StoreHouseStatus(HouseStatus *entity.HouseStatus) (*entity.HouseStatus, []error) {
	ing, errs := HsS.HouseStatusRepo.StoreHouseStatus(HouseStatus)
	if len(errs) > 0 {
		return nil, errs
	}
	return ing, errs
}
