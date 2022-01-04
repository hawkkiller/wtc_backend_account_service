package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main/internal"
	"main/internal/data/model"
	"main/pkg"
	"net/http"
)

func LoginIntoProfile(e echo.Context) error {
	user := new(model.UserProfile)
	userDB := new(model.UserProfile)
	if err := e.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if res := internal.DB.Where("email = ?", user.Email).First(&userDB); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusBadRequest, gorm.ErrRecordNotFound.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, res.Error.Error())

	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err == nil {
			accessToken, refreshToken := pkg.GetTokens(userDB.ID)
			res := make(map[string]interface{})
			res["accessToken"] = accessToken
			res["refreshToken"] = refreshToken

			return e.JSON(http.StatusOK, res)
		} else {
			res := make(map[string]interface{})
			res["message"] = "Password is not correct."
			return e.JSON(http.StatusOK, res)
		}

	}

}
