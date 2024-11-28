package routes

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/wakar473/Ecomerce-Website/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/singup", controllers.Singup())
	incomingRoutes.POST("/users/login", controllers.login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SerachProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())

}
