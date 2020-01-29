package entity

import "time"

// HouseCatagoryMock mocks Food Menu HouseCatagory
var HouseCatagoryMock = HouseCatagory{
	ID:          1,
	Name:        "Mock HouseCatagory 01",
	Description: "Mock HouseCatagory 01 Description",
	Image:       "mock_cat.png",
	HouseTypes:       []HouseType{},
}

// RoleMock mocks Tourist role entity
var RoleMock = Role{
	ID:    1,
	Name:  "Mock Role 01",
}

// HouseTypeMock mocks food menu HouseTypes
var HouseTypeMock = HouseType{
	ID:          1,
	Name:        "Mock HouseType 01",
	Price:       50.5,
	Description: "Mock HouseType 01 Description",
	Categories:  []HouseCatagory{},
	Image:       "mock_HouseType.png",
	HouseStatuss: []HouseStatus{},
}

// HouseStatusMock mocks HouseStatuss in a food HouseType
var HouseStatusMock = HouseStatus{
	ID:          1,
	Name:        "Mock HouseStatus 01",
	Description: "Mock HouseStatus 01 Description",
}

// ReservationMock mocks customer Reservation
var ReservationMock = Reservation{
	ID:        1,
	CreatedAt: time.Time,
	UserID:    1,
	HouseTypeID:    1,
	Quantity:  100,
}

// UserMock mocks application Tourist
var UserMock = Tourist{
	ID:       1,
	FullName: "Mock Tourist 01",
	Email:    "mockTourist@example.com",
	Phone:    "0900000000",
	Password: "P@$$w0rd",
	RoleID:   1,
	Reservations:   []Reservation{},
}

// SessionMock mocks sessions of loged in Tourist
var SessionMock = Session{
	ID:         1,
	UUID:       "_session_one",
	SigningKey: []byte("RestaurantApp"),
	Expires:    0,
}

// FeedbackMock mocks Feedbacks forwarded by application Tourists
var FeedbackMock = Feedback{
	ID:        1,
	FullName:  "Mock Tourist 01",
	Message:   "Mock message",
	Phone:     "0900000000",
	Email:     "mockTourist@example.com",
	CreatedAt: time.Time{},
}
