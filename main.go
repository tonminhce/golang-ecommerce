package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tonminhce/golang-ecommerce/controllers"
	"github.com/tonminhce/golang-ecommerce/database"
	"github.com/tonminhce/golang-ecommerce/middlewares"
	"github.com/tonminhce/golang-ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "products"), database.UserData(database.Client, "users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	routes.Use(middlewares.Authentication())

	router.GET("/additem", app.AddItem())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/checkout", app.Checkout())

	log.Fatal(router.Run(":" + port))
}
