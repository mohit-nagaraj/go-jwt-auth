package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohit-nagaraj/go-jwt-auth/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
}
