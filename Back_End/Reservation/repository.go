package Reservation

import "github.com/gplus1/HouseRentalSystem/Back_End/entity"

// ReservationRepository specifies customer menu Reservation related database operations
type ReservationRepository interface {
	Reservations() ([]entity.Reservation, []error)
	Reservation(id uint) (*entity.Reservation, []error)
	CustomerReservations(customer *Tourist.Tourist) ([]entity.Reservation, []error)
	UpdateReservation(Reservation *entity.Reservation) (*entity.Reservation, []error)
	DeleteReservation(id uint) (*entity.Reservation, []error)
	StoreReservation(Reservation *entity.Reservation) (*entity.Reservation, []error)
}
