package entity

type Userplace struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name      string `json:"name"`
	Kecamatan uint64 `json:"Kecamatan"`
	Provinsi  uint64 `json:"provinsi"`
}
