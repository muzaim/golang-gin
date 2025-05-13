package handler

import (
	"golang-gin/dto"
	"golang-gin/errorhandler"
	"golang-gin/helper"
	"golang-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type studentHandler struct {
	service service.StudentService
}

func NewStudentHandler(s service.StudentService) *studentHandler {
	return &studentHandler{
		service: s,
	}
}

func (h *studentHandler) InsertStudent(c *gin.Context) {
	var student dto.StudentRequest

	if err := c.ShouldBindJSON(&student); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.InsertStudent(&student); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Student Insert Successfully",
	})

	c.JSON(http.StatusCreated, res)
}
