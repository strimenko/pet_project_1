package main

import (
	"pet_project_1/database"
	"pet_project_1/handlers"
	"pet_project_1/models"
	"pet_project_1/repositories"
	"pet_project_1/services"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := database.Init()
	defer db.Close()

	db.AutoMigrate(&models.User{})

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.PUT("/user/:id", userHandler.UpdateUser)
	r.DELETE("/user/:id", userHandler.DeleteUser)
	r.GET("/users", userHandler.GetAllUsers)

	r.Run(":8080")
}
