package handlers

import (
	"lab-seguridad/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Monthly_Taxes(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.UserJWT)
	name := claims.Username
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
