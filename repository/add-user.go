package repository

import (
	"Get_place/entity"
	"gorm.io/gorm"
)

type UserPlaceRepository interface {
	InsertPlace(b entity.Userplace) entity.Userplace
}

type userPlaceConnection struct {
	connection *gorm.DB
}

func NewUserPlaceConnection(connection *gorm.DB) *userPlaceConnection {
	return &userPlaceConnection{connection: connection}
}

func (db userPlaceConnection) InsertPlace(b entity.Userplace) entity.Userplace {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}
