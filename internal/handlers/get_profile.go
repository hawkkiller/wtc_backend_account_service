package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"main/internal"
	"main/internal/data/model"
	"main/pkg"
	"net/http"
)

func GetProfile(e echo.Context) error {
	username := e.QueryParams().Get("username")
	user := model.UserProfile{}
	if err := internal.DB.Where("username = ?", username).Last(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusBadRequest, gorm.ErrRecordNotFound.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if token, err := pkg.NewJWT(model.MapArg{Key: "username", Value: user.ID}); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else {
		res := make(map[string]interface{})
		res["token"] = token
		return e.JSON(http.StatusOK, res)
	}
}
