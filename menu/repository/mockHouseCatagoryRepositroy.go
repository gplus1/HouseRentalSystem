package repository

import (
	"errors"

	"github.com/gplus1/HouseRentalSystem/entity"
	"github.com/gplus1/HouseRentalSystem/menu"
	"github.com/jinzhu/gorm"
)

// MockHouseCatagoryRepo implements the menu.HouseCatagoryRepository interface
type MockHouseCatagoryRepo struct {
	conn *gorm.DB
}

// NewMockHouseCatagoryRepo will create a new object of MockHouseCatagoryRepo
func NewMockHouseCatagoryRepo(db *gorm.DB) menu.HouseCatagoryRepository {
	return &MockHouseCatagoryRepo{conn: db}
}

// Categories returns all fake categories
func (mCatRepo *MockHouseCatagoryRepo) Categories() ([]entity.HouseCatagory, []error) {
	ctgs := []entity.HouseCatagory{entity.HouseCatagoryMock}
	return ctgs, nil
}

// HouseCatagory retrieve a fake HouseCatagory with id 1
func (mCatRepo *MockHouseCatagoryRepo) HouseCatagory(id uint) (*entity.HouseCatagory, []error) {
	ctg := entity.HouseCatagoryMock
	if id == 1 {
		return &ctg, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateHouseCatagory updates a given fake HouseCatagory
func (mCatRepo *MockHouseCatagoryRepo) UpdateHouseCatagory(HouseCatagory *entity.HouseCatagory) (*entity.HouseCatagory, []error) {
	cat := entity.HouseCatagoryMock
	return &cat, nil
}

// DeleteHouseCatagory deletes a given HouseCatagory from the database
func (mCatRepo *MockHouseCatagoryRepo) DeleteHouseCatagory(id uint) (*entity.HouseCatagory, []error) {
	cat := entity.HouseCatagoryMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &cat, nil
}

// StoreHouseCatagory stores a given mock HouseCatagory
func (mCatRepo *MockHouseCatagoryRepo) StoreHouseCatagory(HouseCatagory *entity.HouseCatagory) (*entity.HouseCatagory, []error) {
	cat := HouseCatagory
	return cat, nil
}

// ItemsInHouseCatagory returns mock food menu items
func (mCatRepo *MockHouseCatagoryRepo) ItemsInHouseCatagory(HouseCatagory *entity.HouseCatagory) ([]entity.Item, []error) {
	items := []entity.Item{entity.ItemMock}
	return items, nil
}
