package main

import (
	"Get_place/config"
	"Get_place/repository"
	"fmt"
	"testing"
)

func TestAllKecamatan(t *testing.T) {
	db = config.SetupDatabaseConnection()
	placeRepository = repository.NewPlaceConnection(db)
	fmt.Println(placeRepository.AllKecamatan())

}
