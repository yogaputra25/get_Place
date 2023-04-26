package repository

import (
	"Get_place/entity"
	"gorm.io/gorm"
)

type PlaceRepository interface {
	AllKecamatan() []entity.Kota
	AllProvinsi() []entity.Province
}

func (p placeConnection) AllProvinsi() []entity.Province {
	var Provinsi []entity.Province
	p.connection.Find(&Provinsi)
	return Provinsi
}

func (p placeConnection) AllKecamatan() []entity.Kota {
	var Kecamatan []entity.Kota
	p.connection.Find(&Kecamatan)
	return Kecamatan
}

type placeConnection struct {
	connection *gorm.DB
}

func NewPlaceConnection(connection *gorm.DB) *placeConnection {
	return &placeConnection{connection: connection}
}
