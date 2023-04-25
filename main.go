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
	db              *gorm.DB = config.SetupDatabaseConnection()
	placeRepository          = repository.NewPlaceConnection(db)
	placeService             = service.NewPlaceService(placeRepository)
	placeController          = controller.NewPlaceController(placeService)
)

func main() {

	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	bookRoutes := r.Group("api/place")
	{
		bookRoutes.GET("/", placeController.All)

	}

	r.Run()
	cek := placeRepository.AllKecamatan()
	fmt.Println(cek)
}
