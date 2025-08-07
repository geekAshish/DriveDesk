package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID      uuid.UUID `json:"engine_id"`
	Dispacement   float64   `json:"dispacement"`
	NoOfCylinders float64   `json:"no_of_cylinders"`
	CarRange      float64   `json:"car_range"`
}

type EngineRequest struct {
	Dispacement   float64 `json:"dispacement"`
	NoOfCylinders float64 `json:"no_of_cylinders"`
	CarRange      float64 `json:"car_range"`
}

func ValidateEngineRequest(engineRequest EngineRequest) error {
	if err := validateDispacement(engineRequest.Dispacement); err != nil {
		return err
	}
	if err := validateNoOfCylinders(engineRequest.NoOfCylinders); err != nil {
		return err
	}
	if err := validateCarRange(engineRequest.CarRange); err != nil {
		return err
	}
	return nil
}

func validateDispacement(dispacement float64) error {
	if dispacement <= 0 {
		return errors.New("dispacement must be greater than zero")
	}
	return nil
}
func validateNoOfCylinders(noOfCylinders float64) error {
	if noOfCylinders <= 0 {
		return errors.New("noOfCylinders must be greater than zero")
	}
	return nil
}
func validateCarRange(carRange float64) error {
	if carRange <= 0 {
		return errors.New("carRange must be greater than zero")
	}
	return nil
}
