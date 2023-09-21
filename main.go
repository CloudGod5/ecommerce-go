package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/justintingley/ecommerce-go/controllers"
	"github.com/justintingley/ecommerce-go/database"
	"github.com/justintingley/ecommerce-go/middleware"
	"github.com/justintingley/ecommerce-go/routes"
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
	router.Use(middleware.Authentication())

	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())

	log.Fatal(router.Run(":" + port))

}
