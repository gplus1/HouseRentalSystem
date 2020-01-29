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
	Image        string        `gorm:"type:varchar(255)"`
	HouseStatuss []HouseStatus   `gorm:"many2many:HouseType_HouseStatuss"`
}

type HouseStatus struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
}

type Reservation struct {
	ID          uint
	PlacedAt    time.Time
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
	PlacedAt time.Time `json:"placedat"`
}

type seach struct{

	ID uint `json:"id"`
	Country  []Countries `gorm:"many2many:HouseType_categories"`

}
type transportation struct{
	ID uint `json:"id"`
	transportType `json:"transportType gorm:"type:varchar(255);not null"`
	price float32 `json:"price"`
}

type Rating struct{
	ID uint `json:"id"`
	rated []HouseStatus   `gorm:"many2many:HouseType_HouseStatuss"`
	rate uint `json:"rate"`
}

type Error struct {
	Code    int
	Message string
}
