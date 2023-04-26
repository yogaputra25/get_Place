package controller

import (
	"Get_place/entity"
	"Get_place/helper"
	"Get_place/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PlaceController interface {
	All(context *gin.Context)
	AllProvinsi(context *gin.Context)
}

type placeController struct {
	placeService service.PlaceService
}

func NewPlaceController(placeService service.PlaceService) *placeController {
	return &placeController{placeService: placeService}
}

func (p placeController) All(context *gin.Context) {
	var place []entity.Kota = p.placeService.All()
	res := helper.BuildResponse(true, "OK", place)
	context.JSON(http.StatusOK, res)
}
func (p placeController) AllProvinsi(context *gin.Context) {
	var place []entity.Province = p.placeService.AllProvinsi()
	res := helper.BuildResponse(true, "OK", place)
	context.JSON(http.StatusOK, res)
}
