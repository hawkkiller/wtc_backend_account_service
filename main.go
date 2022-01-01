package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"main/internal"
	"main/internal/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	var err error
	godotenv.Load(".dev.env")

	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Kiev", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	internal.SetDB(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: "mihael", TokenLookup: "header:" + echo.HeaderAuthorization}))

	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("", handlers.GetProfile)
	e.POST("/", handlers.CreateProfile)

	go func() {
		err := e.Start(":9000")
		if err != nil {
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	e.Shutdown(ctx)
}
