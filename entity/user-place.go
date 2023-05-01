package entity

type Userplace struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"column:name"json:"name"`
	Kecamatan uint64 `gorm:"column:kecamatan" json:"Kecamatan"`
	Provinsi  uint64 `gorm:"column:provinsi" json:"provinsi"`
	//Province  Province `json:"province"`
}

type Getalluserplace struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"column:name"json:"name"`
	City_name string `gorm:"column:city_name" json:"Kecamatan"`
	Prov_name string `gorm:"column:prov_name" json:"provinsi"`
}
