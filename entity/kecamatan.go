package entity

type Kota struct {
	City_id   uint64 `json:"id"`
	City_name string `json:"kecamatan"`
	Prov_id   uint64 `json:"idProvinsi"`
}
