package repository

import (
	"github.com/gplus1/HouseRentalSystem//entity"
	"github.com/gplus1/HouseRentalSystem//entity/Tourists/LandLord"
	"github.com/jinzhu/gorm"
)

type LandLordGormRepo struct {
	conn *gorm.DB
}

func NewLandLordGormRepo(db *gorm.DB) LandLord.LandLordRepository {
	return &LandLordGormRepo{conn: db}
}

func (LandLordRepo *LandLordGormRepo) LandLords() ([]entity.LandLord, []error) {
	LandLords := []entity.LandLord{}
	errs := LandLordRepo.conn.Find(&LandLords).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return LandLords, errs
}

func (LandLordRepo *LandLordGormRepo) LandLord(id uint) (*entity.LandLord, []error) {
	LandLord := entity.LandLord{}
	errs := LandLordRepo.conn.First(&LandLord, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &LandLord, errs
}

func (LandLordRepo *LandLordGormRepo) UpdateLandLord(LandLord *entity.LandLord) (*entity.LandLord, []error) {
	usr := LandLord
	errs := LandLordRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (LandLordRepo *LandLordGormRepo) DeleteLandLord(id uint) (*entity.LandLord, []error) {
	usr, errs := LandLordRepo.LandLord(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = LandLordRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (LandLordRepo *LandLordGormRepo) StoreLandLord(LandLord *entity.LandLord) (*entity.LandLord, []error) {
	usr := LandLord
	errs := LandLordRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
