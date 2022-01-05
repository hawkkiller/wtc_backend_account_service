package handlers

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"main/internal"
	"main/internal/data/model"
	"net/http"
)

// RegisterProfile godoc
// @Summary register new profile in WTC system
// @Description register new profile to be able to use WTC.
// @Tags Account
// @Accept json
// @Produce json
// @Param CreateProfileModel body model.RegProfileRequest true "user reg model"
// @Success 200 {object} model.RegProfileResponseOK
// @Failure 400 {object} model.RegProfileResponseBR
// @Failure 403 {object} model.RegProfileResponseFN
// @Router /register [post]
func RegisterProfile(e echo.Context) error {
	user := new(model.UserProfile)
	if err := e.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := e.Validate(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 2)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user.Password = string(password)

	if err := internal.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	return e.JSON(http.StatusOK, model.RegProfileResponseOK{Message: "Successfully registered"})
}
