package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/AtaAksoy/se4458-go-auth-service/docs"
	"github.com/AtaAksoy/se4458-go-auth-service/internal/handler"
	"github.com/AtaAksoy/se4458-go-auth-service/internal/model"
	"github.com/AtaAksoy/se4458-go-auth-service/internal/repository"
	"github.com/AtaAksoy/se4458-go-auth-service/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title Auth Service API
// @version 1.0
// @description Simple Auth Service with Register and Login
// @host localhost:8081
// @BasePath /
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB bağlantı hatası:", err)
	}

	db.AutoMigrate(&model.User{})

	userRepo := repository.NewUserRepository(db)
	authService := &service.AuthService{UserRepo: userRepo}
	authHandler := &handler.AuthHandler{AuthService: authService}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/register", func(c *gin.Context) { authHandler.Register(c.Writer, c.Request) })
	r.POST("/login", func(c *gin.Context) { authHandler.Login(c.Writer, c.Request) })

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server started at :8081")
	r.Run(":8081")
}
