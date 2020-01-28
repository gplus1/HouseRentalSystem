package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/rate"
)

// RateGormRepo implements rate.RateRepository  interface
type RateGormRepo struct {
	Conn *gorm.DB
}
var dbconnect *gorm.DB

// NewRateGormRepo will create a new object of RateGormRepo
func NewRateGormRepo(db *gorm.DB) rate.RateRepository {
	return &RateGormRepo{Conn: db}
}
// RetrieveHotelRateValue retrieve a rate value from the database by its id
func (rRepo *RateGormRepo) RetrieveHotelRateValue(id uint) (*entity.RatingApi, []error) {
	ctg := entity.RatingApi{}
	errs := rRepo.Conn.First(&ctg, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &ctg, errs
}

// AddRate updates a given  in the database
func (rRepo *RateGormRepo) AddRate(updaterate *entity.RatingApi) (*entity.RatingApi, []error) {
	ratevalue := updaterate
	fmt.Println("repository start")
	errs := rRepo.Conn.Create(ratevalue).GetErrors()
	//fmt.Println(errs,1000101001)
	if len(errs) > 0 {
		fmt.Println(errs)
	}
	fmt.Println("repository end")

	return ratevalue, errs
}

func (rRepo *RateGormRepo) GetUserId(updaterate *entity.RatingApi) ([]entity.RatingApi,[]error){
	sval:=[]entity.RatingApi{}
	_,err:=rRepo.Conn.Model(&updaterate).Find(&sval).Rows()
	if err!=nil{ panic(err)}
	return  sval,nil

}
func (rRepo *RateGormRepo) GetUserRateValue(user_id uint) (*entity.RatingApi,[]error){
	data:= entity.RatingApi{}
	err:=rRepo.Conn.Model(&data).Where("user_id=$1",user_id).Find(&data).GetErrors()
	if len(err)>0{
		return  nil,err
	}
	return  &data,nil
}

func (rRepo *RateGormRepo) UpdateUserRateValue(user_id uint,u int) ([]error) {
	user:=&entity.RatingApi{UserId:user_id,RateValue:u}
	errs:=rRepo.Conn.Model(user).Where("user_id=?",user.UserId).UpdateColumns(
		map[string]interface{}{
			"UserId": user.UserId,
			"RateValue": user.RateValue,
		})
	if errs.Error!=nil{
		panic(errs)
	}
	return nil
}


























































