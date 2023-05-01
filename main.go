package main

import (
	"Get_place/config"
	"Get_place/controller"
	"Get_place/repository"
	"Get_place/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                  *gorm.DB = config.SetupDatabaseConnection()
	placeRepository              = repository.NewPlaceConnection(db)
	placeService                 = service.NewPlaceService(placeRepository)
	placeController              = controller.NewPlaceController(placeService)
	UserPlaceRepository          = repository.NewUserPlaceConnection(db)
	UserPlaceService             = service.NewUserPlaceService(UserPlaceRepository)
	userPlaceController          = controller.NewUserPlaceController(UserPlaceService)
)

func main() {

	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	placeRoutes := r.Group("api/place")
	{
		placeRoutes.GET("/kecamatan", placeController.All)
		placeRoutes.GET("/provinsi", placeController.AllProvinsi)

	}

	userPlace := r.Group("api/userplace")
	{
		userPlace.GET("/user", userPlaceController.Get)
		userPlace.POST("/user", userPlaceController.Insert)
		userPlace.DELETE("/user/:id/place", userPlaceController.Delete)
		userPlace.PUT("/user/:id/place", userPlaceController.Update)

	}

	r.Run()
	cek := placeRepository.AllKecamatan()
	fmt.Println(cek)
}
