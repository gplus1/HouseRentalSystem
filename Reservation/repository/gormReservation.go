package repository

import (
	"github.com/gplus1/HouseRentalSystem//entity"
	"github.com/gplus1/HouseRentalSystem//Reservation"
	"github.com/gplus1/jinzhu/gorm"
)

type ReservationGormRepo struct {
	conn *gorm.DB
}

/ NewReservationGormRepo returns new object of ReservationGormRepo
func NewReservationGormRepo(db *gorm.DB) Reservation.ReservationRepository {
	return &ReservationGormRepo{conn: db}
}

/ Reservations returns all customer Reservations stored in the database
func (ReservationRepo *ReservationGormRepo) Reservations() ([]entity.Reservation, []error) {
	Reservations := []entity.Reservation{}
	errs := ReservationRepo.conn.Find(&Reservations).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return Reservations, errs
}

/ Reservation retrieve customer Reservation by Reservation id
func (ReservationRepo *ReservationGormRepo) Reservation(id uint) (*entity.Reservation, []error) {
	Reservation := entity.Reservation{}
	errs := ReservationRepo.conn.First(&Reservation, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &Reservation, errs
}

/ UpdateReservation updates a given customer Reservation in the database
func (ReservationRepo *ReservationGormRepo) UpdateReservation(Reservation *entity.Reservation) (*entity.Reservation, []error) {
	ordr := Reservation
	errs := ReservationRepo.conn.Save(ordr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ordr, errs
}

/ DeleteReservation deletes a given Reservation from the database
func (ReservationRepo *ReservationGormRepo) DeleteReservation(id uint) (*entity.Reservation, []error) {
	ordr, errs := ReservationRepo.Reservation(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = ReservationRepo.conn.Delete(ordr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ordr, errs
}

/ StoreReservation stores a given Reservation in the database
func (ReservationRepo *ReservationGormRepo) StoreReservation(Reservation *entity.Reservation) (*entity.Reservation, []error) {
	ordr := Reservation
	errs := ReservationRepo.conn.Create(ordr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ordr, errs
}

/ CustomerReservations returns list of Reservations from the database for a given customer
func (ReservationRepo *ReservationGormRepo) CustomerReservations(customer *Tourist.Tourist) ([]entity.Reservation, []error) {
	custReservations := []entity.Reservation{}
	errs := ReservationRepo.conn.Model(customer).Related(&custReservations, "Reservations").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return custReservations, errs
}
