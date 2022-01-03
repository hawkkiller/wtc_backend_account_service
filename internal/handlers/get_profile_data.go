package handlers

import (
	"github.com/labstack/echo/v4"
	"main/internal"
	"main/internal/data/model"
	"net/http"
)

func GetProfileData(e echo.Context) error {
	id := e.Get("id").(string)
	userDB := new(model.UserProfile)

	if res := internal.DB.Where("id = ?", id).First(&userDB); res.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, res.Error.Error())
	}

	return e.JSON(http.StatusOK, userDB.ProfileToDB())
}
