package entity

import "time"

type HouseCatagory struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
	Image       string      `gorm:"type:varchar(255)"`
	HouseTypes  []HouseType `gorm:"many2many:HouseType_categories"`
}

type Role struct {
	ID   uint
	Name string `gorm:"type:varchar(255)"`
}

type HouseType struct {
	ID           uint
	Name         string `gorm:"type:varchar(255);not null"`
	Price        float32
	Description  string
	Categories   []HouseCatagory `gorm:"many2many:HouseType_categories"`
	Image        string          `gorm:"type:varchar(255)"`
	HouseStatuss []HouseStatus   `gorm:"many2many:HouseType_HouseStatuss"`
}

type HouseStatus struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
}

type Reservation struct {
	ID          uint
	time    time.Time
	UserID      uint
	HouseTypeID uint
	Quantity    uint
}

type Feedback struct {
	ID       uint      `json:"id"`
	FullName string    `json:"fullname" gorm:"type:varchar(255)"`
	Message  string    `json:"message"`
	Phone    string    `json:"phone" gorm:"type:varchar(100);not null; unique"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null; unique"`
	time time.Time `json:"placedat"`
}

type Error struct {
	Code    int
	Message string
}
