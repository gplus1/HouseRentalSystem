package users

import (
	"time"
	"github.com/HouseRentalSystem/entity"
)
type Tourist struct {
	
	ID       uint
	UserName string `gorm:"type:varchar(100);not null"`
	FullName string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);not null; unique"`
	Country string `gorm:"type:varchar(100);not null"`
	State string `gorm:"type:varchar(100);not null"`
	City string `gorm:"type:varchar(100);not null"`
	Phone    string `gorm:"type:varchar(100);not null; unique"`
    Atendee uint
    Childs uint 
	Password string `gorm:"type:varchar(100)"`
	Roles    []Role `gorm:"many2many:user_roles"`
	Reservation   []Reservation
}