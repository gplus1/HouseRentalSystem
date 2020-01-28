package repository

import (
	"github.com/gplus1/HouseRentalSystem//entity"
	"github.com/gplus1/HouseRentalSystem//feedback"
	"github.com/jinzhu/gorm"
)

type FeedBackGormRepo struct {
	conn *gorm.DB
}

func NewFeedBackGormRepo(db *gorm.DB) FeedBack.FeedBackRepository {
	return &FeedBack_Gorm_Repo{conn: db}
}

func (cmntRepo *FeedBack_Gorm_Repo) FeedBack() ([]entity.FeedBack, []error) {
	cmnts := []entity.FeedBack{}
	errs := cmntRepo.conn.Find(&cmnts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

func (cmntRepo *FeedBack_Gorm_Repo) FeedBack(id uint) (*entity.FeedBack, []error) {
	cmnt := entity.FeedBack{}
	errs := cmntRepo.conn.First(&cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cmnt, errs
}
func (cmntRepo *FeedBack_Gorm_Repo) UpdateFeedBack(FeedBack *entity.FeedBack) (*entity.FeedBack, []error) {
	cmnt := FeedBack
	errs := cmntRepo.conn.Save(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
func (cmntRepo *FeedBack_Gorm_Repo) DeleteFeedBack(id uint) (*entity.FeedBack, []error) {
	cmnt, errs := cmntRepo.FeedBack(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cmntRepo.conn.Delete(cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
func (cmntRepo *FeedBack_Gorm_Repo) StoreFeedBack(FeedBack *entity.FeedBack) (*entity.FeedBack, []error) {
	cmnt := FeedBack
	errs := cmntRepo.conn.Create(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}