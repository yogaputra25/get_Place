package dto

type UserPlaceDTO struct {
	Name      string `json:"name"`
	Kecamatan uint64 `json:"kecamatan"`
	Provinsi  uint64 `json:"provinsi"`
}
