package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohit-nagaraj/go-jwt-auth/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users")
	incomingRoutes.GET("/users/:user_id")
}
