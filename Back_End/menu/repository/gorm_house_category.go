package repository

import (
	"github.com/HouseRentalSystem/Back_End/entity"
	"github.com/HouseRentalSystem/Back_End/entity/menu"
	"github.com/jinzhu/gorm"
)

type HCategoryGormRepo struct {
	conn *gorm.DB
}

func NewHCategoryGormRepo(db *gorm.DB) menu.CategoryRepository {
	return &HCategoryGormRepo{conn: db}
}

func (cRepo *HCategoryGormRepo) Categories() ([]entity.House_category, []error) {
	ctgs := []entity.House_category{}
	errs := cRepo.conn.Find(&ctgs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ctgs, errs
}
func (cRepo *HCategoryGormRepo) Category(id uint) (*entity.House_category, []error) {
	ctg := entity.House_category{}
	errs := cRepo.conn.First(&ctg, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &ctg, errs
}

func (cRepo *HCategoryGormRepo) UpdateCategory(category *entity.House_category) (*entity.House_category, []error) {
	cat := category
	errs := cRepo.conn.Save(cat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}

func (cRepo *HCategoryGormRepo) DeleteCategory(id uint) (*entity.House_category, []error) {
	cat, errs := cRepo.Category(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = cRepo.conn.Delete(cat, cat.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}
func (cRepo *HCategoryGormRepo) StoreCategory(category *entity.House_category) (*entity.House_category, []error) {
	cat := category
	errs := cRepo.conn.Create(cat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}

func (cRepo *HCategoryGormRepo) HouseTypesInCategory(category *entity.House_category) ([]entity.HouseType, []error) {
	HouseTypes := []entity.HouseType{}
	cat, errs := cRepo.Category(category.ID)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cRepo.conn.Model(cat).Related(&HouseTypes, "HouseTypes").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return HouseTypes, errs
}
