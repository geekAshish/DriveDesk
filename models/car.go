package models

import (
	"errors"
	"slices"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Year     string    `json:"year"`
	Brand    string    `json:"brand"`
	FuelType string    `json:"fuel_type"`
	Price    float64   `json:"price"`
	Engine   Engine    `json:"engine"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	Name     string    `json:"name"`
	Year     string    `json:"year"`
	Brand    string    `json:"brand"`
	FuelType string    `json:"fuel_type"`
	Price    float64   `json:"price"`
	Engine   Engine    `json:"engine"`
}

func ValidateRequest(carRequest CarRequest) error {
	if err := validateName(carRequest.Name); err != nil {
		return err
	}
	if err := validateYear(carRequest.Year); err != nil {
		return err
	}
	if err := validateFuelType(carRequest.FuelType); err != nil {
		return err
	}
	if err := validateEngine(carRequest.Engine); err != nil {
		return err
	}
	if err := validatePrice(carRequest.Price); err != nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}

	return  nil
}

func validateYear(year string) error {
	if year == "" {
		return errors.New("year is required")
	}

	_, err := strconv.Atoi(year);
	if err != nil {
		return errors.New("year must be a valid number")
	}

	currentYear := time.Now().Year()
	yearInt, _ := strconv.Atoi(year)
	if yearInt < 1886 || yearInt > currentYear {
		return errors.New("year must be between 1886 and the current year")
	}

	return  nil
}


func validateFuelType(fuelType string) error {
	if fuelType == "" {
		return errors.New("fuel type is required")
	}

	validFuelType := []string{"petrol", "diesel", "electric", "hybrid"};

	isValid := slices.Contains(validFuelType, fuelType)

	if isValid {
		return  nil
	}

	return errors.New("not a valid fuel type")
}


func validateEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil {
		return errors.New("engine ID is required")
	}

	if engine.Dispacement <= 0 {
		return errors.New("engine displacement must be greater than zero")
	}

	if engine.NoOfCylinders <= 0 {
		return errors.New("number of cylinders must be greater than zero")
	}

	if engine.CarRange <= 0 {
		return errors.New("car range must be greater than zero")
	}

	return nil
}

func validatePrice(price float64) error {
	if price <= 0 {
		return errors.New("price must be greater than zero")
	}

	return nil
}

 