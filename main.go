package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	_ "main/docs/wtc"
	"main/internal"
	"main/internal/handlers"
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
// @version 0.4+040120221154
// @description Account service for WTC.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email miskadl09@gmail.com

// @host localhost:9000
// @BasePath /api/v1
// @schemes http
func main() {
	var err error
	godotenv.Load(".dev.env")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Kiev", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DBNAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	internal.SetDB(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &CustomValidator{validator: validator.New()}

	api := e.Group("/api/v1")

	api.GET("/swagger/*", echoSwagger.WrapHandler)

	api.POST("/data", handlers.GetProfileData, middlewares.CheckJWT())

	api.POST("/login", handlers.LoginIntoProfile)

	api.POST("/register", handlers.RegisterProfile)

	go func() {
		err := e.Start(":9000")
		if err != nil {
			log.Println(err)

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
