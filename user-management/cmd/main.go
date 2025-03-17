package main

import (
	"fmt"
	"user-management/delivery/http"
	"user-management/infrastructure"
	"user-management/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi repository (In-Memory)
	userRepo := infrastructure.NewInMemoryUserRepo()

	// Inisialisasi use case
	userUseCase := usecase.NewUserUseCase(userRepo)
	authUseCase := usecase.NewAuthUseCase(userRepo)

	// Inisialisasi controller
	userController := http.NewUserController(userUseCase, authUseCase)

	// Inisialisasi router Gin
	r := gin.Default()

	// Routing API
	api := r.Group("/api")
	{
		api.POST("/register", userController.RegisterUser)
		api.POST("/login", userController.Login)

		protected := api.Group("/")
		protected.Use(userController.AuthMiddleware)
		{
			protected.GET("/profile", userController.GetProfile)
			protected.POST("/logout", userController.Logout)
		}
	}

	// Menjalankan server
	port := "8080"
	fmt.Println("Server berjalan di port " + port)
	r.Run(":" + port)
}
