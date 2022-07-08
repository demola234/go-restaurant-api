package main

import (
	"go-restaurant/database"
	middleware "go-restaurant/middleware"
	routers "go-restaurant/routers"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routers.UserRoutes(router)
	router.Use(middleware.Authentication())
	routers.FoodRoutes(router)
	routers.TableRoutes(router)
	routers.MenuRoutes(router)
	routers.OrderRoutes(router)
	routers.OrderItemRoutes(router)
	routers.InvoiceRoutes(router)

	router.Run(":" + port)
}
