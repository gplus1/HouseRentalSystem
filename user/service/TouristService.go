package service

import (
	"github.com/gplus1/HouseRentalSystem//entity/Tourists"
	"github.com/gplus1/HouseRentalSystem//entity"
)


type TouristService struct {
	TouristRepo Tourist.TouristRepository
}


func NewTouristService(TouristRepository Tourist.TouristRepository) Tourist.TouristService {
	return &TouristService{TouristRepo: TouristRepository}
}


func (Ts *TouristService) Tourists() ([]Tourist.Tourist, []error) {
	Tsrs, errs := Ts.TouristRepo.Tourists()
	if len(errs) > 0 {
		return nil, errs
	}
	return Tsrs, errs
}


func (Ts *TouristService) Tourist(id uint) (*Tourist.Tourist, []error) {
	Tsr, errs := Ts.TouristRepo.Tourist(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return Tsr, errs
}


func (Ts *TouristService) UpdateTourist(Tourist *Tourist.Tourist) (*Tourist.Tourist, []error) {
	Tsr, errs := Ts.TouristRepo.UpdateTourist(Tourist)
	if len(errs) > 0 {
		return nil, errs
	}
	return Tsr, errs
}


func (Ts *TouristService) DeleteTourist(id uint) (*Tourist.Tourist, []error) {
	Tsr, errs := Ts.TouristRepo.DeleteTourist(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return Tsr, errs
}


func (Ts *TouristService) StoreTourist(Tourist *Tourist.Tourist) (*Tourist.Tourist, []error) {
	Tsr, errs := Ts.TouristRepo.StoreTourist(Tourist)
	if len(errs) > 0 {
		return nil, errs
	}
	return Tsr, errs
}
