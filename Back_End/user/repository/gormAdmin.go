package repository

import (
	"github.com/gplus1/betsegawlemma/restaurant/Admin"
	"github.com/gplus1/betsegawlemma/restaurant/entity"
	"github.com/gplus1/jinzhu/gorm"
)

type AdminGormRepo struct {
	conn *gorm.DB
}

func NewAdminGormRepo(db *gorm.DB) Admin.AdminRepository {
	return &AdminGormRepo{conn: db}
}

func (AdminRepo *AdminGormRepo) Admins() ([]entity.Admin, []error) {
	Admins := []entity.Admin{}
	errs := AdminRepo.conn.Find(&Admins).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return Admins, errs
}

func (AdminRepo *AdminGormRepo) Admin(id uint) (*entity.Admin, []error) {
	Admin := entity.Admin{}
	errs := AdminRepo.conn.First(&Admin, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &Admin, errs
}

func (AdminRepo *AdminGormRepo) UpdateAdmin(Admin *entity.Admin) (*entity.Admin, []error) {
	usr := Admin
	errs := AdminRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (AdminRepo *AdminGormRepo) DeleteAdmin(id uint) (*entity.Admin, []error) {
	usr, errs := AdminRepo.Admin(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = AdminRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (AdminRepo *AdminGormRepo) StoreAdmin(Admin *entity.Admin) (*entity.Admin, []error) {
	usr := Admin
	errs := AdminRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
