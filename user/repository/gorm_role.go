package repository

import (
	"github.com/gplus1/HouseRentalSystem//entity"
	"github.com/gplus1/HouseRentalSystem//Tourist"
	"github.com/jinzhu/gorm"
)

type RoleGormRepo struct {
	conn *gorm.DB
}

func NewRoleGormRepo(db *gorm.DB) Tourist.RoleRepository {
	return &RoleGormRepo{conn: db}
}

func (roleRepo *RoleGormRepo) Roles() ([]entity.Role, []error) {
	roles := []entity.Role{}
	errs := roleRepo.conn.Find(&roles).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return roles, errs
}

func (roleRepo *RoleGormRepo) Role(id uint) (*entity.Role, []error) {
	role := entity.Role{}
	errs := roleRepo.conn.First(&role, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &role, errs
}

func (roleRepo *RoleGormRepo) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	r := role
	errs := roleRepo.conn.Save(r).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}

func (roleRepo *RoleGormRepo) DeleteRole(id uint) (*entity.Role, []error) {
	r, errs := roleRepo.Role(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = roleRepo.conn.Delete(r, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}

func (roleRepo *RoleGormRepo) StoreRole(role *entity.Role) (*entity.Role, []error) {
	r := role
	errs := roleRepo.conn.Create(r).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}
