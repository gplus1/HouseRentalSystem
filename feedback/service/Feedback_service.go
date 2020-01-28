package service

import (
	"github.com/HouseRentalSystem//feedback"
	"github.com/HouseRentalSystem//entity"

)

/ FeedBackService implements menu.FeedBackService interface
type FeedBackService struct {
	FeedBackRepo FeedBack.FeedBackRepository
}

/ NewFeedBackService returns a new FeedBackService object
func NewFeedBackService(commRepo FeedBack.FeedBackRepository) FeedBack.FeedBackService {
	return &FeedBackService{FeedBackRepo: commRepo}
}

/ FeedBacks returns all stored FeedBacks
func (cs *FeedBackService) FeedBacks() ([]entity.FeedBack, []error) {
	cmnts, errs := cs.FeedBackRepo.FeedBacks()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

/ FeedBack retrieves stored FeedBack by its id
func (cs *FeedBackService) FeedBack(id uint) (*entity.FeedBack, []error) {
	cmnt, errs := cs.FeedBackRepo.FeedBack(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

/ UpdateFeedBack updates a given FeedBack
func (cs *FeedBackService) UpdateFeedBack(FeedBack *entity.FeedBack) (*entity.FeedBack, []error) {
	cmnt, errs := cs.FeedBackRepo.UpdateFeedBack(FeedBack)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

/ DeleteFeedBack deletes a given FeedBack
func (cs *FeedBackService) DeleteFeedBack(id uint) (*entity.FeedBack, []error) {
	cmnt, errs := cs.FeedBackRepo.DeleteFeedBack(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

/ StoreFeedBack stores a given FeedBack
func (cs *FeedBackService) StoreFeedBack(FeedBack *entity.FeedBack) (*entity.FeedBack, []error) {
	cmnt, errs := cs.FeedBackRepo.StoreFeedBack(FeedBack)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
