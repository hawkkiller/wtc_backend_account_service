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
	"strconv"
	"time"
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
			accessToken, _ := pkg.NewJWT(time.Duration(24)*time.Hour, strconv.Itoa(int(userDB.ID)))
			refreshToken, _ := pkg.NewJWT(time.Duration(24)*time.Hour*30, strconv.Itoa(int(userDB.ID)))

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
