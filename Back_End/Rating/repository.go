package rate

import "github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"

type RateRepository interface {

	//rate methods

	AddRate(updaterate *entity.RatingApi) (*entity.RatingApi, []error)
	RetrieveHotelRateValue(id uint) (*entity.RatingApi, []error)
	GetUserId(updaterate *entity.RatingApi) ([]entity.RatingApi, []error)
	GetUserRateValue(user_id uint) (*entity.RatingApi, []error)
	UpdateUserRateValue(user_id uint, u int) ([]error)
	
}