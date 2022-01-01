package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"main/internal"
	"main/internal/data/model"
	"net/http"
)

func GetProfile(e echo.Context) error {
	id := e.QueryParams().Get("username")
	fmt.Println(id, "username")
	user := model.UserProfile{}
	internal.DB.Where("username = ?", id).First(&user)
	fmt.Println(user, "user")

	return e.JSON(http.StatusOK, user)
}
