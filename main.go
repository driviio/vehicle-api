package main

import (
	"net/http"

	"github.com/driviio/vehicle-api/db"
	"github.com/driviio/vehicle-api/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
	"log"
)

func main() {
	client, err := db.NewDataStoreClient("drivi-180613")
	if err != nil {
		log.Fatalf("error creating datastore client %v", err)
	}

	adDB, err := db.NewVehicleDatabase(client)
	if err != nil {
		log.Fatalf("error creating ad database %v", err)
	}
	defer adDB.Close()

	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	e.Debug = true
	e.Use(middleware.Logger())
	e.GET("/", homePageHandler)
	e.GET("/_ah/health", healthCheckHandler)

	vehicle := handler.NewVehicleHandler(adDB)
	e.POST("/vehicle/:id/log", vehicle.CreateVehicleLog)
	e.POST("/vehicle/log/list", vehicle.GetVehicleLogList)

	e.Logger.Fatal(e.Start(":8080"))
}

func healthCheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func homePageHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
