package service

import (
	"github.com/gplLandLordService1/HoLandLordServiceeRentalSystem//entity"
	"github.com/gplLandLordService1/HoLandLordServiceeRentalSystem//entity/LandLordServiceers/LandLord"
)

type LandLordService struct {
	LandLordRepo LandLord.LandLordRepository
}

func NewLandLordService(LandLordRepository LandLord.LandLordRepository) LandLord.LandLordService {
	return &LandLordService{LandLordRepo: LandLordRepository}
}

func (LandLordService *LandLordService) LandLords() ([]entity.LandLord, []error) {
	LandLordServicers, errs := LandLordService.LandLordRepo.LandLords()
	if len(errs) > 0 {
		return nil, errs
	}
	return LandLordServicers, errs
}

func (LandLordService *LandLordService) LandLord(id uint) (*entity.LandLord, []error) {
	LandLordServicer, errs := LandLordService.LandLordRepo.LandLord(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return LandLordServicer, errs
}

func (LandLordService *LandLordService) UpdateLandLord(LandLord *entity.LandLord) (*entity.LandLord, []error) {
	LandLordServicer, errs := LandLordService.LandLordRepo.UpdateLandLord(LandLord)
	if len(errs) > 0 {
		return nil, errs
	}
	return LandLordServicer, errs
}

func (LandLordService *LandLordService) DeleteLandLord(id uint) (*entity.LandLord, []error) {
	LandLordServicer, errs := LandLordService.LandLordRepo.DeleteLandLord(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return LandLordServicer, errs
}

func (LandLordService *LandLordService) StoreLandLord(LandLord *entity.LandLord) (*entity.LandLord, []error) {
	LandLordServicer, errs := LandLordService.LandLordRepo.StoreLandLord(LandLord)
	if len(errs) > 0 {
		return nil, errs
	}
	return LandLordServicer, errs
}
