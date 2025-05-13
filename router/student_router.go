package router

import (
	"golang-gin/config"
	"golang-gin/handler"
	"golang-gin/repository"
	"golang-gin/service"

	"github.com/gin-gonic/gin"
)

func StudentRouter(api *gin.RouterGroup) {

	studentRepository := repository.NewStudentRepository(config.DB)
	studentService := service.NewStudentService(studentRepository)
	studentHandler := handler.NewStudentHandler(studentService)

	api.POST("/insert", studentHandler.InsertStudent)
}
