package Reservation

import (
	"github.com/gplus1/HouseRentalSystem/entity"
	"github.com/gplus1/HouseRentalSystem/entity/"
)

type ReservationRepository interface{
	Reservations() ([]entity.Reservation, []error)
	Reservation(id uint) (*entity.Reservation, []error)
	CustomerReservations(customer *user.Tourist) ([]entity.Reservation, []error)
	UpdateReservation(Reservation *entity.Reservation) (*entity.Reservation, []error)
	DeleteReservation(id uint) (*entity.Reservation, []error)
	StoreReservation(Reservation *entity.Reservation) (*entity.Reservation, []error)
}
