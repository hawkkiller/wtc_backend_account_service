package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/happierall/l"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	_ "main/docs/wtc"
	"main/internal"
	"main/internal/handlers/user"
	"main/pkg/middlewares"
	"net/http"
	"os"
	"os/signal"
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

// @title WTC ACCOUNT SERVICE
// @version 0.8+07-01-2022-22:30
// @description Account service for WTC.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email miskadl09@gmail.com

// @host localhost:9000
// @BasePath /api/v1
// @schemes http
func main() {
	var err error
	err = godotenv.Load(".dev.env")
	if err != nil {
		log.Fatalf("Variable environment not found")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DBNAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	internal.SetDB(db)

	e := echo.New()
	e.Use(middlewares.Logger())
	e.Use(middleware.Recover())

	e.Validator = &CustomValidator{validator: validator.New()}

	api := e.Group("/api/v1")

	api.GET("/swagger/*", echoSwagger.WrapHandler)

	api.GET("/data", user.GetProfileData, middlewares.CheckJWT("Authorization"))

	api.PUT("/update", user.UpdateProfile, middlewares.CheckJWT("Authorization"))

	api.GET("/refresh", user.Refresh, middlewares.CheckJWT("Refresh"))

	api.POST("/login", user.Login)

	api.POST("/register", user.Register)

	go func() {
		err := e.Start(":9000")
		if err != nil {
			l.Print(err)

			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

}
