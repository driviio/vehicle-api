package model

type VehicleLog struct {
	VehicleID      int64  `json:"vehicleId"`
	Data   string `json:"data" validate:"required"`
}

