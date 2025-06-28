package routes

import (
	controller "go_jwt_auth_gin_gonic/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.SignUp())
	incomingRoutes.POST("users/login", controller.Login())
}
