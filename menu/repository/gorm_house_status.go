package repository

import (
	"github.com/gplus1/HouseRentalSystem//entity"
	"github.com/gplus1/HouseRentalSystem//entity/menu"
	"github.com/jinzhu/gorm"
)

/house status
type HouseStatusGormRepo struct {
	conn *gorm.DB
}

/creating new house status
func NewHouseStatusGormRepo(db *gorm.DB) menu.HouseStatusRepository {
	return &HouseStatusGormRepo{conn: db}
}

/house status
func (HsRepo *HouseStatusGormRepo) HouseStatuss() ([]entity.HouseStatus, []error) {
	ingts := []entity.HouseStatus{}
	errs := HsRepo.conn.Find(&ingts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ingts, errs
}

/house status
func (HsRepo *HouseStatusGormRepo) HouseStatus(id uint) (*entity.HouseStatus, []error) {
	ingt := entity.HouseStatus{}
	errs := HsRepo.conn.First(&ingt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &ingt, errs
}
/update status
func (HsRepo *HouseStatusGormRepo) UpdateHouseStatus(HouseStatus *entity.HouseStatus) (*entity.HouseStatus, []error) {
	ingt := HouseStatus
	errs := HsRepo.conn.Save(ingt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ingt, errs
}
/delete status
func (HsRepo *HouseStatusGormRepo) DeleteHouseStatus(id uint) (*entity.HouseStatus, []error) {
	ingt, errs := HsRepo.HouseStatus(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = HsRepo.conn.Delete(ingt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ingt, errs
}
/store status
func (HsRepo *HouseStatusGormRepo) StoreHouseStatus(HouseStatus *entity.HouseStatus) (*entity.HouseStatus, []error) {
	ingt := HouseStatus
	errs := HsRepo.conn.Create(ingt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ingt, errs
}
