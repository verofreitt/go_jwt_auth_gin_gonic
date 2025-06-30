package routes

import (
	"go_jwt_auth_gin_gonic/middleware"

	controller "go_jwt_auth_gin_gonic/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// Define user-related routes here
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
