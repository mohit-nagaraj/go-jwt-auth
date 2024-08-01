package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohit-nagaraj/go-jwt-auth/controllers"
	"github.com/mohit-nagaraj/go-jwt-auth/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
