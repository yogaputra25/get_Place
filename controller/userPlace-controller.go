package controller

import (
	"Get_place/dto"
	"Get_place/entity"
	"Get_place/helper"
	"Get_place/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserPlaceController interface {
	Insert(ctx *gin.Context)
	Delete(ctx *gin.Context)
	UPdate(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type userPlaceController struct {
	userPlaceService service.UserPlaceService
}

func NewUserPlaceController(userPlaceService service.UserPlaceService) *userPlaceController {
	return &userPlaceController{userPlaceService: userPlaceService}
}
func (c userPlaceController) Delete(ctx *gin.Context) {
	var userPlace entity.Userplace

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		userPlace.ID = id
		c.userPlaceService.Delete(userPlace)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	}

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

func (c userPlaceController) Update(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
	var userPlace dto.UserPlaceUpdateDTO
	errDTO := ctx.ShouldBind(&userPlace)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {

		//result := c.userPlaceService.Update(userPlace, id)
		c.userPlaceService.Update(userPlace, id)
		response := helper.BuildResponse(true, "OK", "data Updated")
		ctx.JSON(http.StatusCreated, response)
	}
}

func (p userPlaceController) Get(context *gin.Context) {
	var userplace []entity.Getalluserplace = p.userPlaceService.Get()
	res := helper.BuildResponse(true, "OK", userplace)
	context.JSON(http.StatusOK, res)
}

//func (c userPlaceController) Update(ctx *gin.Context) {
//
//	var userPlace dto.UserPlaceUpdateDTO
//
//	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
//	if err != nil {
//		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
//		ctx.JSON(http.StatusBadRequest, response)
//	} else {
//		userPlace.Id = id
//		c.userPlaceService.Update(userPlace)
//		res := helper.BuildResponse(true, "Updated", helper.EmptyObj{})
//		ctx.JSON(http.StatusOK, res)
//	}
//}
