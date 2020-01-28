package menu

import "github.com/gplus1/HouseRentalSystem//entity"

type HouseCatagoryService interface {
	Categories() ([]entity.HouseCatagory, []error)
	HouseCatagory(id uint) (*entity.HouseCatagory, []error)
	UpdateHouseCatagory(HouseCatagory *entity.HouseCatagory) (*entity.HouseCatagory, []error)
	DeleteHouseCatagory(id uint) (*entity.HouseCatagory, []error)
	StoreHouseCatagory(HouseCatagory *entity.HouseCatagory) (*entity.HouseCatagory, []error)
	HouseTypesInHouseCatagory(HouseCatagory *entity.HouseCatagory) ([]entity.HouseType, []error)
}

type HouseTypeService interface {
	HouseTypes() ([]entity.HouseType, []error)
	HouseType(id uint) (*entity.HouseType, []error)
	UpdateHouseType(menu *entity.HouseType) (*entity.HouseType, []error)
	DeleteHouseType(id uint) (*entity.HouseType, []error)
	StoreHouseType(HouseType *entity.HouseType) (*entity.HouseType, []error)
}

type HouseStatusService interface {
	HouseStatuss() ([]entity.HouseStatus, []error)
	HouseStatus(id uint) (*entity.HouseStatus, []error)
	UpdateHouseStatus(HouseStatus *entity.HouseStatus) (*entity.HouseStatus, []error)
	DeleteHouseStatus(id uint) (*entity.HouseStatus, []error)
	StoreHouseStatus(HouseStatus *entity.HouseStatus) (*entity.HouseStatus, []error)
}
