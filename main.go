package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RianIhsan/go-auth-rbac/controller"
	"github.com/RianIhsan/go-auth-rbac/db"
	"github.com/RianIhsan/go-auth-rbac/model"
	"github.com/RianIhsan/go-auth-rbac/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	loadDatabase()

	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error load .env file")
	}
}

func loadDatabase() {
	db.InitDB()
	db.Db.AutoMigrate(&model.Role{})
	db.Db.AutoMigrate(&model.User{})
	seedData()
}

func serveApplication() {
	router := gin.Default()

	authRoutes := router.Group("/auth/user")
	authRoutes.POST("/register", controller.Register)
	authRoutes.POST("/login", controller.Login)

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(util.JWTAuth())
	adminRoutes.GET("/users", controller.GetUsers)
	adminRoutes.GET("/user/:id", controller.GetUser)
	adminRoutes.PUT("/user/:id", controller.UpdateUser)
	adminRoutes.POST("/user/role", controller.CreateRole)
	adminRoutes.GET("/user/roles", controller.GetRoles)
	adminRoutes.PUT("/user/role/:id", controller.UpdateRole)
	adminRoutes.POST("/room/add", controller.CreateRoom)
	adminRoutes.PUT("/room/:id", controller.UpdateRoom)
	adminRoutes.GET("/room/bookings", controller.GetBookings)

	publicRoutes := router.Group("/api/view")
	publicRoutes.GET("/rooms", controller.GetRooms)
	publicRoutes.GET("/room/:id", controller.GetRoom)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(util.JWTAuthCustomer())
	protectedRoutes.GET("/rooms/booked", controller.GetUserBookings)
	protectedRoutes.POST("/room/book", controller.CreateBooking)

	router.Run(":8000")
	fmt.Println("Server is running on port 8000")
}

func seedData() {
	var roles = []model.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "anonymous", Description: "Unauthenticated customer role"}}
	var user = []model.User{{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1}}
	db.Db.Save(&roles)
	db.Db.Save(&user)
}
