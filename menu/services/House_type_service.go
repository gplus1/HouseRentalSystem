package service

import (
	"github.com/gplus1/HouseRentalSystem//entity"
	"github.com/gplus1/HouseRentalSystem//entity/menu"
)

type HTypeservice struct {
	HouseTypeRepo menu.HouseTypeRepository
}

func NewHTypeservice(HouseTypeRepository menu.HouseTypeRepository) menu.HTypeservice {
	return &HTypeservice{HouseTypeRepo: HouseTypeRepository}
}

func (HtS *HTypeservice) HouseTypes() ([]entity.HouseType, []error) {
	itms, errs := HtS.HouseTypeRepo.HouseTypes()
	if len(errs) > 0 {
		return nil, errs
	}
	return itms, errs
}

func (HtS *HTypeservice) HouseType(id uint) (*entity.HouseType, []error) {
	itm, errs := HtS.HouseTypeRepo.HouseType(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

func (HtS *HTypeservice) UpdateHouseType(HouseType *entity.HouseType) (*entity.HouseType, []error) {
	itm, errs := HtS.HouseTypeRepo.UpdateHouseType(HouseType)
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

func (HtS *HTypeservice) DeleteHouseType(id uint) (*entity.HouseType, []error) {
	itm, errs := HtS.HouseTypeRepo.DeleteHouseType(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

func (HtS *HTypeservice) StoreHouseType(HouseType *entity.HouseType) (*entity.HouseType, []error) {
	itm, errs := HtS.HouseTypeRepo.StoreHouseType(HouseType)
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}
