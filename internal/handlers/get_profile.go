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
	email := e.QueryParams().Get("email")
	user := model.UserProfile{}
	if err := internal.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusBadRequest, gorm.ErrRecordNotFound.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if token, err := pkg.NewJWT(model.MapArg{Key: "id", Value: user.ID}); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else {
		res := make(map[string]interface{})
		res["token"] = token
		return e.JSON(http.StatusOK, res)
	}
}
