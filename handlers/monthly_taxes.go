package handlers

import (
	"fmt"
	"io"
	"lab-seguridad/models"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Monthly_Taxes(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.UserJWT)
	name := claims.Username

	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("sunat"), nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create request")
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", name))

	resp, err := client.Do(req)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get response")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.String(http.StatusBadGateway, "Failed to fetch data")
	}

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to read response data")
	}

	return c.Blob(http.StatusOK, resp.Header.Get("Content-Type"), rawData)

}
