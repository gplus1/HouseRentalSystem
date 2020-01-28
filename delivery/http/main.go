package main

import (
	"net/http"

	"github.com/gplus1/HouseRentalSystem/FeedBack/repository"
	"github.com/gplus1/HouseRentalSystem//FeedBack/service"
	"github.com/gplus1/HouseRentalSystem//delivery/http/handler"
	urepim "github.com/gplus1/HouseRentalSystem//Tourist/repository"
	usrvim "github.com/gplus1/HouseRentalSystem//Tourist/service"
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
	adminRoleHandler := handler.NewLandLordRoleHandler(roleSrv)

	FeedBackRepo := repository.FeedBackGormRepo(dbconn)
	FeedBackrv := service.NewFeedBackervice(FeedBackRepo)
	adminFeedBackHandler := handler.NewLandLordFeedBackHandler(FeedBackrv)

	router := httprouter.New()

	router.GET("/v1/LandLord/roles", adminRoleHandler.GetRoles)

	router.GET("/v1/LandLord/FeedBack/:id", adminFeedBackHandler.GetSingleFeedBack)
	router.GET("/v1/LandLord/FeedBack", adminFeedBackHandler.GetFeedBack)
	router.PUT("/v1/LandLord/FeedBack/:id", adminFeedBackHandler.PutFeedBack)
	router.POST("/v1/LandLord/FeedBack", adminFeedBackHandler.PostFeedBack)
	router.DELETE("/v1/LandLord/FeedBack/:id", adminFeedBackHandler.DeleteFeedBack)

	http.ListenAndServe(":8181", router)
}
