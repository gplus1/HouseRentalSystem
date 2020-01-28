package users

import (
	"time"
	"github.com/HouseRentalSystem/Back_End/entity"
)
type Admin struct {
	gorm.Model
	ID       uint
	UserName string `gorm:"type:varchar(255);not null"`
	FullName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null; unique"`
	Phone    string `gorm:"type:varchar(100);not null; unique"`
	Password string `gorm:"type:varchar(255)"`
	Roles    []Role `gorm:"many2many:user_roles"`
	Reservations   []Reservation
}