package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mpavithran/go-crud/config"
	"github.com/mpavithran/go-crud/controllers"
	"github.com/mpavithran/go-crud/repositories"
	"github.com/mpavithran/go-crud/routes"
	"github.com/mpavithran/go-crud/services"
)

func main() {
	config.ConnectDB()

	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r := gin.Default()

	routes.UserRoutes(r, userController)

	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}
