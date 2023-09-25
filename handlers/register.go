package handlers

import (
	"lab-seguridad/models"
	"lab-seguridad/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var user_body models.User
	err := c.Bind(&user_body)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user_body.Password = util.Hash_Password(user_body.Password)

	result := util.DB_Connection.Create(user_body)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, "try a different username")
	}

	return c.String(http.StatusAccepted, "success")
}
