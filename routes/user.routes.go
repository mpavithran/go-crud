package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mpavithran/go-crud/controllers"
)

func UserRoutes(router *gin.Engine, userController *controllers.UserController) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", userController.CreateUser)
		userGroup.GET("/", userController.GetUsers)
		userGroup.GET("/:id", userController.GetUserByID)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}
}
