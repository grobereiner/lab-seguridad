package main

import (
	// "gorm.io/driver/sqlite"

	"lab-seguridad/handlers"
	"lab-seguridad/models"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup_Server() *echo.Echo {
	e := echo.New()
	e.POST("/login", handlers.Login)
	e.POST("/register", handlers.Register)
	return e
}

func Connect_Database() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("dsn")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
	return db
}
