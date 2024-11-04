package main

import (
	"gin_REST_API_ex/src/config"
	"gin_REST_API_ex/src/controller"
	"gin_REST_API_ex/src/middleware"
	"gin_REST_API_ex/src/model"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")

	if err := config.MakeConnectionMongoDB(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}

	user := model.User{ID: uint(0), Username: "raven", Password: "testRaven621"}
	user.BeforeSave()

	_, err := config.MongoClient.Database(os.Getenv("DB_NAME")).Collection("user").InsertOne(config.CTX, user)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := gin.Default()

	publicRoutes := router.Group("/api")
	{
		publicRoutes.POST("/login", controller.Login) // Authenticate a user and provide a token
	}

	protectedRoutes := router.Group("/api/admin")
	{
		protectedRoutes.Use(middleware.JwtAuthMiddleware())
		protectedRoutes.POST("/users", controller.CreateUser)           // Create a new user
		protectedRoutes.GET("/users/me", controller.GetCurrentAuthUser) // Get the current authenticated user
		protectedRoutes.POST("/logout", controller.Logout)              // Invalidate the user token
		protectedRoutes.GET("/users/:id", controller.GetUserByID)       // Get a user if current user is authenticated user
		protectedRoutes.GET("/users", controller.GetUsers)              // Get all users if current user is authenticated user
	}

	router.Run("localhost:8080")

}
