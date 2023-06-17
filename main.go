// (c) Copyright 2023, Shaunak Sensarma.

package main

import (
	"log"
	"os"

	"github.com/ShaunakSensarma/GolangEcommerce/controllers"
	"github.com/ShaunakSensarma/GolangEcommerce/database"
	"github.com/ShaunakSensarma/GolangEcommerce/middleware"
	"github.com/ShaunakSensarma/GolangEcommerce/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
