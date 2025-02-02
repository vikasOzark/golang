package handlers

import (
	"echo/api"
	"echo/database/connection"
	"echo/database/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Users(c echo.Context) error {
	var user []models.User

	result := connection.Connect().Find(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	userResult := api.NewResponse[[]models.User, any](true, "User list fetched successfully.", http.StatusOK, user, nil)
	return api.ResponseProvider(userResult, c)
}
