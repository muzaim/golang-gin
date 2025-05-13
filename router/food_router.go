package router

import (
	"golang-gin/config"
	"golang-gin/handler"
	"golang-gin/repository"
	"golang-gin/service"

	"github.com/gin-gonic/gin"
)

func FoodRouter(api *gin.RouterGroup) {
	foodRepository := repository.NewFoodRepository(config.DB)
	foodService := service.NewFoodService(foodRepository)
	foodHandler := handler.NewFoodHandler(foodService)

	api.POST("/food", foodHandler.CreateFood)
}
