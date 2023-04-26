package main

import (
	"Get_place/config"
	"Get_place/controller"
	"Get_place/dto"
	"Get_place/repository"
	"Get_place/service"
	"fmt"
	"testing"
)

func TestAllKecamatan(t *testing.T) {
	db = config.SetupDatabaseConnection()
	placeRepository = repository.NewPlaceConnection(db)
	placeService = service.NewPlaceService(placeRepository)
	UserPlaceRepository = repository.NewUserPlaceConnection(db)
	UserPlaceService = service.NewUserPlaceService(UserPlaceRepository)
	userPlaceController = controller.NewUserPlaceController(UserPlaceService)
	//fmt.Println(placeRepository.AllProvinsi())
	//fmt.Println(placeRepository.AllKecamatan())
	//fmt.Println(placeService.All())
	fmt.Println(UserPlaceService.InsertUser(dto.UserPlaceDTO{Name: "komsd", Kecamatan: 1, Provinsi: 2}))

}
