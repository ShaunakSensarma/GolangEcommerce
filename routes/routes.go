// (c) Copyright 2023, Shaunak Sensarma.

package routes

import (
	"github.com/ShaunakSensarma/GolangEcommerce/controllers"
	"github.com/gin-gonic/gin"
)

//UserRoutes handles the incoming Postman API requests.
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
