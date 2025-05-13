package main

import (
	"fmt"
	"golang-gin/config"
	"golang-gin/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()

	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.AuthRouter(api)
	router.FoodRouter(api)
	router.StudentRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
}
