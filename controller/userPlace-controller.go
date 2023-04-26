package controller

import (
	"Get_place/dto"
	"Get_place/helper"
	"Get_place/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserPlaceController interface {
	Insert(ctx *gin.Context)
}

type userPlaceController struct {
	userPlaceService service.UserPlaceService
}

func NewUserPlaceController(userPlaceService service.UserPlaceService) *userPlaceController {
	return &userPlaceController{userPlaceService: userPlaceService}
}

func (c userPlaceController) Insert(ctx *gin.Context) {
	var userPlace dto.UserPlaceDTO
	errDTO := ctx.ShouldBind(&userPlace)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {

		result := c.userPlaceService.InsertUser(userPlace)
		response := helper.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusCreated, response)
	}
}
