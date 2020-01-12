package service

import (
	"github.com/gplAdminService1/HoAdminServiceeRentalSystem/Back_End/entity"
	"github.com/gplAdminService1/HoAdminServiceeRentalSystem/Back_End/entity/AdminServiceers/Admin"
)

type AdminService struct {
	AdminRepo Admin.AdminRepository
}

func NewAdminService(AdminRepository Admin.AdminRepository) Admin.AdminService {
	return &AdminService{AdminRepo: AdminRepository}
}

func (AdminService *AdminService) Admins() ([]entity.Admin, []error) {
	AdminServicers, errs := AdminService.AdminRepo.Admins()
	if len(errs) > 0 {
		return nil, errs
	}
	return AdminServicers, errs
}

func (AdminService *AdminService) Admin(id uint) (*entity.Admin, []error) {
	AdminServicer, errs := AdminService.AdminRepo.Admin(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return AdminServicer, errs
}

func (AdminService *AdminService) UpdateAdmin(Admin *entity.Admin) (*entity.Admin, []error) {
	AdminServicer, errs := AdminService.AdminRepo.UpdateAdmin(Admin)
	if len(errs) > 0 {
		return nil, errs
	}
	return AdminServicer, errs
}

func (AdminService *AdminService) DeleteAdmin(id uint) (*entity.Admin, []error) {
	AdminServicer, errs := AdminService.AdminRepo.DeleteAdmin(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return AdminServicer, errs
}

func (AdminService *AdminService) StoreAdmin(Admin *entity.Admin) (*entity.Admin, []error) {
	AdminServicer, errs := AdminService.AdminRepo.StoreAdmin(Admin)
	if len(errs) > 0 {
		return nil, errs
	}
	return AdminServicer, errs
}
