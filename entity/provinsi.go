package entity

type Province struct {
	Prov_id   uint64 ` gorm:"column:prov_id" json:"id"`
	Prov_name string `gorm:"column:prov_name" json:"provName"`
}
