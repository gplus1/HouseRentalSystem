package main

import (
	"net/http"

	"github.com/gplus1/HouseRentalSystem/Back_End/FeedBack/repository"
	"github.com/gplus1/HouseRentalSystem/Back_End/FeedBack/service"
	"github.com/gplus1/HouseRentalSystem/Back_End/delivery/http/handler"
	urepim "github.com/gplus1/HouseRentalSystem/Back_End/user/repository"
	usrvim "github.com/gplus1/HouseRentalSystem/Back_End/user/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

func main() {

	dbconn, err := gorm.Open("postgres",
		"postgres:/postgres:P@$$w0rdD2@localhost/restaurantdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	roleRepo := urepim.NewRoleGormRepo(dbconn)
	roleSrv := usrvim.NewRoleService(roleRepo)
	adminRoleHandler := handler.NewAdminRoleHandler(roleSrv)

	FeedBackRepo := repository.FeedBackGormRepo(dbconn)
	FeedBackrv := service.NewFeedBackervice(FeedBackRepo)
	adminFeedBackHandler := handler.NewAdminFeedBackHandler(FeedBackrv)

	router := httprouter.New()

	router.GET("/v1/admin/roles", adminRoleHandler.GetRoles)

	router.GET("/v1/admin/FeedBack/:id", adminFeedBackHandler.GetSingleFeedBack)
	router.GET("/v1/admin/FeedBack", adminFeedBackHandler.GetFeedBack)
	router.PUT("/v1/admin/FeedBack/:id", adminFeedBackHandler.PutFeedBack)
	router.POST("/v1/admin/FeedBack", adminFeedBackHandler.PostFeedBack)
	router.DELETE("/v1/admin/FeedBack/:id", adminFeedBackHandler.DeleteFeedBack)

	http.ListenAndServe(":8181", router)
}
