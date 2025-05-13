package handler

import (
	"golang-gin/dto"
	"golang-gin/errorhandler"
	"golang-gin/helper"
	"golang-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type foodHandler struct {
	service service.FoodService
}

func NewFoodHandler(s service.FoodService) *foodHandler {
	return &foodHandler{
		service: s,
	}
}

func (h *foodHandler) CreateFood(c *gin.Context) {
	var food dto.FoodRequest

	if err := c.ShouldBindJSON(&food); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.CreateFood(&food); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Food Created Successfully",
	})

	c.JSON(http.StatusCreated, res)
}
