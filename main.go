package main

import (
	"log"
	"os"

	"github.com/CloudGod5/ecommerce-go/controllers"
	"github.com/CloudGod5/ecommerce-go/database"
	"github.com/CloudGod5/ecommerce-go/middleware"
	"github.com/CloudGod5/ecommerce-go/routes"
	"github.com/gin-gonic/gin"
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

	router.GET("/addtocart", app.AddToCart)
	router.GET("/removeitem", app.RemoveItem)
	router.GET("/cartcheckout", app.BuyFromCart)
	router.GET("/instantbuy", app.InstantBuy)

	log.Fatal(router.Run(":" + port))

}
