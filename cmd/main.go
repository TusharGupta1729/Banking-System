package main

import (
	"fmt"
	"os"

	"banking-system/config"
	"banking-system/handlers"
	"banking-system/models"
	"banking-system/repository"
	"banking-system/routes"
	"banking-system/services"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	if err := config.ConnectDatabase(); err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}

	fmt.Println("Database connected successfully")

	//-------------------------------------------------------------------------------------------------------
	// DATABSE CONNECTED WITH SERVER
	//-------------------------------------------------------------------------------------------------------

	err := config.DB.AutoMigrate(
		&models.Bank{},
		&models.Branch{},
		&models.Customer{},
		&models.Account{},
		&models.Transaction{},
		&models.Loan{},
	)

	if err != nil {
		fmt.Println("Migration failed:", err)
		return
	}

	fmt.Println("Migration completed successfully")

	//-------------------------------------------------------------------------------------------------------
	//-------------------------------------------------------------------------------------------------------

	gin.SetMode(gin.ReleaseMode)

	// Creating gin router
	r := gin.Default()
	if err := r.SetTrustedProxies(nil); err != nil {
		fmt.Println("Failed to configure trusted proxies:", err)
		return
	}

	repo := repository.NewBankRepository()

	service := services.NewBankService(repo)

	handler := handlers.NewBankHandler(service)

	routes.SetupRoutes(r, handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	if err := r.Run(":" + port); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
