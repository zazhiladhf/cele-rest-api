package routes

import (
	"project-rest-api/auth"
	"project-rest-api/handler"
	"project-rest-api/middleware"
	"project-rest-api/user"

	"github.com/gin-gonic/gin"
)

func MasterAuthorRoutes(api *gin.RouterGroup, authService auth.Service, userService user.Service, handler *handler.MasterAuthorHandler) {
	api.GET("/author", handler.Get)
	api.POST("/author/store", middleware.AuthMiddleware(authService, userService), handler.Create)
	api.PUT("/author/update/:id", middleware.AuthMiddleware(authService, userService), handler.Update)
}
