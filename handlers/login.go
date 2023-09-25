package handlers

import (
	"lab-seguridad/models"
	"lab-seguridad/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var user_body models.User
	err := c.Bind(&user_body)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user_body.Password = util.Hash_Password(user_body.Password)

	var user_query models.User
	result := util.DB_Connection.Where(&models.User{Username: user_body.Username}).Take(&user_query)

	if result.Error != nil || user_query.Password != user_body.Password {
		return c.String(http.StatusBadRequest, "bad username or password")
	}

	return c.String(http.StatusAccepted, "success")
}
