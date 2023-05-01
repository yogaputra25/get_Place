package repository

import (
	"Get_place/entity"
	"gorm.io/gorm"
)

type UserPlaceRepository interface {
	InsertPlace(b entity.Userplace) entity.Userplace
	Delete(b entity.Userplace)
	Update(b entity.Userplace, id uint64) entity.Userplace
	Get() []entity.Getalluserplace
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
func (db userPlaceConnection) Update(b entity.Userplace, id uint64) entity.Userplace {
	//var userPlace entity.Userplace
	//db.connection.First(&userPlace, &b.ID)
	db.connection.Where("id = ?", id).Updates(entity.Userplace{Name: b.Name, Kecamatan: b.Kecamatan, Provinsi: b.Provinsi})
	//db.connection.Find(&userPlace, b.ID)
	//db.connection.Save(&userPlace)
	//db.connection.Find(&b)
	return b
}

func (db userPlaceConnection) Delete(b entity.Userplace) {
	db.connection.Delete(&b)
}
func (db userPlaceConnection) Get() []entity.Getalluserplace {
	var b []entity.Getalluserplace

	db.connection.Table("userplace").Select("userplace.id, userplace.name,kota.city_name,province.prov_name").
		Joins("LEFT JOIN province ON userplace.provinsi = province.prov_id").
		Joins("LEFT JOIN kota ON userplace.kecamatan = kota.city_id").
		Scan(&b)
	return b
}
