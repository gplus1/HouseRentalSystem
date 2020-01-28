package repository

import (
	"github.com/gplus1/HouseRentalSystem//Tourist"
	"github.com/gplus1/HouseRentalSystem//Tourist"
	"github.com/jinzhu/gorm"
)

type TouristGormRepo struct {
	conn *gorm.DB
}

func NewTouristGormRepo(db *gorm.DB) Tourist.TouristRepository {
	return &TouristGormRepo{conn: db}
}

func (TouristRepo *TouristGormRepo) Tourists() ([]entity.Tourist, []error) {
	Tourists := []entity.Tourist{}
	errs := TouristRepo.conn.Find(&Tourists).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return Tourists, errs
}

func (TouristRepo *TouristGormRepo) Tourist(id uint) (*entity.Tourist, []error) {
	Tourist := entity.Tourist{}
	errs := TouristRepo.conn.First(&Tourist, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &Tourist, errs
}

func (TouristRepo *TouristGormRepo) UpdateTourist(Tourist *entity.Tourist) (*entity.Tourist, []error) {
	usr := Tourist
	errs := TouristRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (TouristRepo *TouristGormRepo) DeleteTourist(id uint) (*entity.Tourist, []error) {
	usr, errs := TouristRepo.Tourist(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = TouristRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (TouristRepo *TouristGormRepo) StoreTourist(Tourist *entity.Tourist) (*entity.Tourist, []error) {
	usr := Tourist
	errs := TouristRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
