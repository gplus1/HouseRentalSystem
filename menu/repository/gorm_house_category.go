package repository

import (
	"github.com/gplus1/HouseRentalSystem//entity"
	"github.com/gplus1/HouseRentalSystem//entity/menu"
	"github.com/jinzhu/gorm"
)

type HCategoryGormRepo struct {
	conn *gorm.DB
}

func NewHCategoryGormRepo(db *gorm.DB) menu.CategoryRepository {
	return &HCategoryGormRepo{conn: db}
}

func (HcRepo *HCategoryGormRepo) Categories() ([]entity.Housecategory, []error) {
	ctgs := []entity.Housecategory{}
	errs := HcRepo.conn.Find(&ctgs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ctgs, errs
}
func (HcRepo *HCategoryGormRepo) Category(id uint) (*entity.House_category, []error) {
	ctg := entity.House_category{}
	errs := HcRepo.conn.First(&ctg, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &ctg, errs
}

func (HcRepo *HCategoryGormRepo) UpdateCategory(category *entity.House_category) (*entity.House_category, []error) {
	cat := category
	errs := HcRepo.conn.Save(cat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}

func (HcRepo *HCategoryGormRepo) DeleteCategory(id uint) (*entity.House_category, []error) {
	cat, errs := HcRepo.Category(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = HcRepo.conn.Delete(cat, cat.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}
func (HcRepo *HCategoryGormRepo) StoreCategory(category *entity.House_category) (*entity.House_category, []error) {
	cat := category
	errs := HcRepo.conn.Create(cat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}

func (HcRepo *HCategoryGormRepo) HouseTypesInCategory(category *entity.House_category) ([]entity.HouseType, []error) {
	HouseTypes := []entity.HouseType{}
	cat, errs := HcRepo.Category(category.ID)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = HcRepo.conn.Model(cat).Related(&HouseTypes, "HouseTypes").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return HouseTypes, errs
}
