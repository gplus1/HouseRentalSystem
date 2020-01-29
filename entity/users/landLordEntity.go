package users

import (
	"time"
	"github.com/HouseRentalSystem/entity/users"
)
type LandLord struct {
	gorm.Model
	ID       uint
	UserName string `gorm:"type:varchar(100);not null"`
	FullName string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);not null; unique"`
	Country string `gorm:"type:varchar(100);not null"`
	State string `gorm:"type:varchar(100);not null"`
	City string `gorm:"type:varchar(100);not null"`
	HouseNo string `gorm:"type:varchar(50); not null; unique"`
	Phone    string `gorm:"type:varchar(100);not null; unique"`
	Password string `gorm:"type:varchar(100)"`
	
}

