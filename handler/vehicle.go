package handler

import (
	"github.com/driviio/vehicle-api/db"
	"github.com/driviio/vehicle-api/model"
	"github.com/labstack/echo"
	"net/http"
)

func NewVehicleHandler(vehicleDB db.VehicleDatabase) *vehicleHandler {
	return &vehicleHandler{vehicleDB: vehicleDB}
}

type vehicleHandler struct {
	vehicleDB db.VehicleDatabase
}

func (h *vehicleHandler) CreateVehicleLog(c echo.Context) error {
	log := new(model.VehicleLog)
	if err := c.Bind(log); err != nil {
		return err
	}
	dbLog := &db.VehicleLog{
		VehicleID:ParseInt64(c.Param("vehicleId")),
		Data:log.Data,
	}
	_, err := h.vehicleDB.AddVehicleLog(dbLog)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h *vehicleHandler) GetVehicleLogList(c echo.Context) error {
	log, err := h.vehicleDB.ListVehicleLog()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, log)
}