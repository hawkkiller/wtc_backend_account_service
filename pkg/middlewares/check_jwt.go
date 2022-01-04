package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
)

func CheckJWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			var id interface{}
			req := c.Request()
			tokenReq := req.Header.Get("Authorization")
			t := strings.Split(tokenReq, " ")[1]
			claims := jwt.MapClaims{}
			_, err = jwt.ParseWithClaims(t, &claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("SECRET")), nil
			})
			if err != nil {
				return err
			}

			for key, val := range claims {
				if key == "jti" {
					id = val
				}
			}

			if id != nil {
				c.Set("id", id)
				return next(c)
			} else {
				return echo.ErrBadGateway
			}

		}
	}
}
