package Feedback

import "github.com/gplus1/HouseRentalSystem//entity"

/ FeedbackRepository specifies customer Feedback related database operations
type FeedbackRepository interface {
	Feedbacks() ([]entity.Feedback, []error)
	Feedback(id uint) (*entity.Feedback, []error)
	UpdateFeedback(Feedback *entity.Feedback) (*entity.Feedback, []error)
	DeleteFeedback(id uint) (*entity.Feedback, []error)
	StoreFeedback(Feedback *entity.Feedback) (*entity.Feedback, []error)
}
