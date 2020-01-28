package repository

import (
	"github.com/gplus1/HouseRentalSystem//entity"
	"github.com/gplus1/HouseRentalSystem//entity/menu"
	"github.com/jinzhu/gorm"
)

type HTypeGormRepo struct {
	conn *gorm.DB
}

func NewHTypeGormRepo(db *gorm.DB) menu.HouseTypeRepository {
	return &HTypeGormRepo{conn: db}
}

func (HouseTypeRepo *HTypeGormRepo) HouseTypes() ([]entity.HouseType, []error) {
	HouseTypes := []entity.HouseType{}
	errs := HouseTypeRepo.conn.Find(&HouseTypes).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return HouseTypes, errs
}

func (HouseTypeRepo *HTypeGormRepo) HouseType(id uint) (*entity.HouseType, []error) {
	HouseType := entity.HouseType{}
	errs := HouseTypeRepo.conn.First(&HouseType, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &HouseType, errs
}

func (HouseTypeRepo *HTypeGormRepo) UpdateHouseType(HouseType *entity.HouseType) (*entity.HouseType, []error) {
	itm := HouseType
	errs := HouseTypeRepo.conn.Save(itm).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

func (HouseTypeRepo *HTypeGormRepo) DeleteHouseType(id uint) (*entity.HouseType, []error) {
	itm, errs := HouseTypeRepo.HouseType(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = HouseTypeRepo.conn.Delete(itm, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

func (HouseTypeRepo *HTypeGormRepo) StoreHouseType(HouseType *entity.HouseType) (*entity.HouseType, []error) {
	itm := HouseType
	errs := HouseTypeRepo.conn.Create(itm).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}
