package service

import (
	"github.com/gplus1/HouseRentalSystem//Entity/Tourists/LandLord"
)

type LandLordService struct {
	LandLordRepo LandLord.LandLordRepository
}

func NewLandLordService(LandLordRepository LandLord.LandLordRepository) LandLord.LandLordService {
	return &LandLordService{LandLordRepo: LandLordRepository}
}

func (Ts *LandLordService) LandLords() ([]LandLord.LandLord, []error) {
	Tsrs, errs := Ts.LandLordRepo.LandLords()
	if len(errs) > 0 {
		return nil, errs
	}
	return Tsrs, errs
}

func (Ts *LandLordService) LandLord(id uint) (*LandLord.LandLord, []error) {
	Tsr, errs := Ts.LandLordRepo.LandLord(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return Tsr, errs
}

func (Ts *LandLordService) UpdateLandLord(LandLord *LandLord.LandLord) (*LandLord.LandLord, []error) {
	Tsr, errs := Ts.LandLordRepo.UpdateLandLord(LandLord)
	if len(errs) > 0 {
		return nil, errs
	}
	return Tsr, errs
}

func (Ts *LandLordService) DeleteLandLord(id uint) (*LandLord.LandLord, []error) {
	Tsr, errs := Ts.LandLordRepo.DeleteLandLord(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return Tsr, errs
}

func (Ts *LandLordService) StoreLandLord(LandLord *LandLord.LandLord) (*LandLord.LandLord, []error) {
	Tsr, errs := Ts.LandLordRepo.StoreLandLord(LandLord)
	if len(errs) > 0 {
		return nil, errs
	}
	return Tsr, errs
}
