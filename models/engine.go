package models

import "github.com/google/uuid"

type Engine struct {
	EngineID      uuid.UUID `json:"engine_id"`
	Dispacement   float64   `json:"dispacement"`
	NoOfCylinders float64   `json:"no_of_cylinders"`
	CarRange	  float64   `json:"car_range"`
}
