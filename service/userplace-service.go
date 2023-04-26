package service

import (
	"Get_place/dto"
	"Get_place/entity"
	"Get_place/repository"
)

type UserPlaceService interface {
	InsertUser(d dto.UserPlaceDTO) entity.Userplace
}

type userPlaceService struct {
	userPlaceRepository repository.UserPlaceRepository
}

func NewUserPlaceService(userPlaceRepository repository.UserPlaceRepository) *userPlaceService {
	return &userPlaceService{userPlaceRepository: userPlaceRepository}
}

func (b userPlaceService) InsertUser(d dto.UserPlaceDTO) entity.Userplace {
	userPlace := entity.Userplace{Name: d.Name, Kecamatan: d.Kecamatan, Provinsi: d.Provinsi}

	return b.userPlaceRepository.InsertPlace(userPlace)

}
