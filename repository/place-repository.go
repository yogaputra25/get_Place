package repository

import (
	"Get_place/entity"
	"gorm.io/gorm"
)

type PlaceRepository interface {
	AllKecamatan() []entity.Kota
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
