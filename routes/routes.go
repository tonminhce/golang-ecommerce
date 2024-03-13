package routes

import (
	"github.com/tonminhce/golang-ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {

}

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/user/register", controllers.Register())
	incomingRoutes.POST("/user/login", controllers.Login())
	incomingRoutes.POST("/admin/adminproduct", controllers.AdminProduct())
	incomingRoutes.GET("/users/products", controllers.GetProducts())
	incomingRoutes.GET("/users/search", controllers.SearchProducts())
}