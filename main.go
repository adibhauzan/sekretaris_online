package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/adibhauzan/sekretaris_online_backend/controllers"
	"github.com/adibhauzan/sekretaris_online_backend/handlers"
	"github.com/adibhauzan/sekretaris_online_backend/models"
	"github.com/adibhauzan/sekretaris_online_backend/repositoryimpl"

	"github.com/adibhauzan/sekretaris_online_backend/routes"
	_ "github.com/gin-gonic/gin"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	db = db.Debug()
	err = db.AutoMigrate(&models.Status{}, &models.Jadwal{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to auto-migrate models: ", err)
	}

	statusRepo := repositoryimpl.NewStatusRepository(db)
	jadwalRepo := repositoryimpl.NewJadwalRepository(db)
	userRepo := repositoryimpl.NewUserRepository(db)

	statusController := controllers.NewStatusController(statusRepo)
	jadwalController := controllers.NewJadwalController(jadwalRepo)
	userController := controllers.NewUserController(userRepo)

	go handlers.AutoDeleteExpiredData(db)

	r := routes.SetupRouter(jadwalController, statusController, userController)

	port := ":8000"

	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start the server: ", err)
	}

	fmt.Printf("Server is running on port %s\n", port)
}
