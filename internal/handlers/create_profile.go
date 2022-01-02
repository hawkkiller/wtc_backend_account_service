package handlers

import (
	"github.com/labstack/echo/v4"
	"main/internal"
	"main/internal/data/model"
	"net/http"
)

func CreateProfile(e echo.Context) error {
	user := new(model.UserProfile)
	if err := e.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := e.Validate(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := internal.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	return e.JSON(http.StatusOK, &user)
}
