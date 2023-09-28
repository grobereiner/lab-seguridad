package main

import (
	"lab-seguridad/handlers"
	"lab-seguridad/models"
	"os"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup_Server() *echo.Echo {
	e := echo.New()

    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
      AllowOrigins: []string{"http://localhost:5173"},
      AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
    })) 

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", handlers.Login)
	e.POST("/register", handlers.Register)

	r := e.Group("/monthly-taxes")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.UserJWT)
		},
		SigningKey: []byte(os.Getenv("jwt")),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("",handlers.Monthly_Taxes)
	// e.GET("/monthly-taxes", handlers.Monthly_Taxes, echojwt.WithConfig(config))

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
