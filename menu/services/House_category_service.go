package service

import (
	"github.com/gplus1/HouseRentalSystem//entity"
	"github.com/gplus1/HouseRentalSystem//entity/menu"
)

/ H_catagory_service implements menu.H_catagory_service interface
type H_catagory_service struct {
	categoryRepo menu.CategoryRepository
}

/ NewH_catagory_service will create new H_catagory_service object
func NewH_catagory_service(CatRepo menu.CategoryRepository) menu.H_catagory_service {
	return &H_catagory_service{categoryRepo: CatRepo}
}

/ Categories returns list of categories
func (HcS *H_catagory_service) Categories() ([]entity.Category, []error) {

	categories, errs := HcS.categoryRepo.Categories()

	if len(errs) > 0 {
		return nil, errs
	}

	return categories, nil
}
func (HcS *H_catagory_service) StoreCategory(category *entity.Category) (*entity.Category, []error) {

	cat, errs := HHcS.categoryRepo.StoreCategory(category)

	if len(errs) > 0 {
		return nil, errs
	}

	return cat, nil
}
func (HcS *H_catagory_service) Category(id uint) (*entity.Category, []error) {

	c, errs := HcS.categoryRepo.Category(id)

	if len(errs) > 0 {
		return c, errs
	}

	return c, nil
}

func (HcS *H_catagory_service) UpdateCategory(category *entity.Category) (*entity.Category, []error) {

	cat, errs := HcS.categoryRepo.UpdateCategory(category)

	if len(errs) > 0 {
		return nil, errs
	}

	return cat, nil
}

func (HcS *H_catagory_service) DeleteCategory(id uint) (*entity.Category, []error) {

	cat, errs := HcS.categoryRepo.DeleteCategory(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return cat, nil
}

func (HcS *H_catagory_service) HouseTypesInCategory(category *entity.Category) ([]entity.HouseType, []error) {

	cts, errs := HcS.categoryRepo.HouseTypesInCategory(category)

	if len(errs) > 0 {
		return nil, errs
	}

	return cts, nil

}
