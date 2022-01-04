package middlewares

import (
	"github.com/happierall/l"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Logger() echo.MiddlewareFunc {
	type Log struct {
		Headers http.Header
		Iat     string
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			log := Log{Headers: req.Header, Iat: time.Now().Format(time.UnixDate)}
			l.Print(log)
			return next(c)
		}
	}
}
