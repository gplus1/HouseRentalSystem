package service

import (
	"github.com/gplus1/HouseRentalSystem/entity"
	"github.com/gplus1/HouseRentalSystem/Reservation"
)

/ ReservationService implements menu.ReservationService interface
type ReservationService struct {
	ReservationRepo Reservation.ReservationRepository
}

/ NewReservationService returns new ReservationService object
func NewReservationService(ReservationRepository Reservation.ReservationRepository) Reservation.ReservationService {
	return &ReservationService{ReservationRepo: ReservationRepository}
}

/ Reservations returns all stored food Reservations
func (os *ReservationService) Reservations() ([]entity.Reservation, []error) {
	ords, errs := os.ReservationRepo.Reservations()
	if len(errs) > 0 {
		return nil, errs
	}
	return ords, errs
}

/ Reservation retrieves an Reservation by its id
func (os *ReservationService) Reservation(id uint) (*entity.Reservation, []error) {
	ord, errs := os.ReservationRepo.Reservation(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}

/ CustomerReservations returns all Reservations of a given customer
func (os *ReservationService) CustomerReservations(customer *Tourist.Tourist) ([]entity.Reservation, []error) {
	ords, errs := os.ReservationRepo.CustomerReservations(customer)
	if len(errs) > 0 {
		return nil, errs
	}
	return ords, errs
}

/ UpdateReservation updates a given Reservation
func (os *ReservationService) UpdateReservation(Reservation *entity.Reservation) (*entity.Reservation, []error) {
	ord, errs := os.ReservationRepo.UpdateReservation(Reservation)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}

/ DeleteReservation deletes a given Reservation
func (os *ReservationService) DeleteReservation(id uint) (*entity.Reservation, []error) {
	ord, errs := os.DeleteReservation(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}

/ StoreReservation stores a given Reservation
func (os *ReservationService) StoreReservation(Reservation *entity.Reservation) (*entity.Reservation, []error) {
	ord, errs := os.StoreReservation(Reservation)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}
