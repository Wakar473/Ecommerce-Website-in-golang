package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wakar473/Ecommerce-Website/controllers"
	"github.com/wakar473/Ecommerce-Website/database"
	"github.com/wakar473/Ecommerce-Website/middleware"
	"github.com/wakar473/Ecommerce-Website/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"

	}
	/*with the help of app varibale we can controll all the routes*/
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.Client, "Users")

	// Create a router
	router := gin.new()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.Addtocart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromcart())
	router.GET("/instantbuy", app.InstantBuy())

}
