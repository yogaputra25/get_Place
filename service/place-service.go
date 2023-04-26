package service

import (
	"Get_place/entity"
	"Get_place/repository"
)

type PlaceService interface {
	All() []entity.Kota
	AllProvinsi() []entity.Province
}

type placeService struct {
	placeRepository repository.PlaceRepository
}

func NewPlaceService(placeRepository repository.PlaceRepository) *placeService {
	return &placeService{placeRepository: placeRepository}
}

func (p placeService) All() []entity.Kota {
	return p.placeRepository.AllKecamatan()
}

func (p placeService) AllProvinsi() []entity.Province {
	return p.placeRepository.AllProvinsi()
}
