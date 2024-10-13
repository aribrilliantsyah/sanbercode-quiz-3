package routes

import (
	"book-management/app/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authController controllers.AuthController
}

func NewRouteAuth(authController controllers.AuthController) AuthRoutes {
	return AuthRoutes{authController}
}

func (r *AuthRoutes) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")
	router.POST("/login", r.authController.Login)
	router.POST("/register", r.authController.Register)
}
