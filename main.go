package main

import (
	"os"

	"github.com/jonathanbs9/go-restaurant/database"
	"github.com/jonathanbs9/go-restaurant/middleware"
	"github.com/jonathanbs9/go-restaurant/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.client, "food")

func main() {
	p := os.Getenv("PORT")
	if p == " " {
		p = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OderRoutes(router)
	routes.OrderItemRoutes(router)

	router.Run(":" + p)

}
