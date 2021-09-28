package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jonathanbs9/go-restaurant/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:foodId", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:foodId", controller.UpdateFood())
}
