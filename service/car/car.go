package car

import (
	"context"

	"github.com/geekAshish/DriveDesk/models"
	"github.com/geekAshish/DriveDesk/store"
	"go.opentelemetry.io/otel"
)

type CarService struct {
	store store.CarStoreInterface
}

func NewCarService(store store.CarStoreInterface) *CarService {
	return &CarService{store: store}
}

func (s *CarService) GetCarById(ctx context.Context, id string) (*models.Car, error) {
	tracer := otel.Tracer("CarService")
	ctx, span := tracer.Start(ctx, "GetCarById-Service")
	defer span.End()

	car, err := s.store.GetCarById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &car, nil
}

func (s *CarService) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	tracer := otel.Tracer("CarService")
	ctx, span := tracer.Start(ctx, "GetCarByBrand-Service")
	defer span.End()

	cars, err := s.store.GetCarByBrand(ctx, brand, isEngine)
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (s *CarService) CreateCar(ctx context.Context, car *models.CarRequest) (*models.Car, error) {
	tracer := otel.Tracer("CarService")
	ctx, span := tracer.Start(ctx, "CreateCar-Service")
	defer span.End()

	if err := models.ValidateRequest(*car); err != nil {
		return nil, err
	}

	createdCar, err := s.store.CreateCar(ctx, car)
	if err != nil {
		return nil, err
	}

	return &createdCar, nil
}

func (s *CarService) UpdateCar(ctx context.Context, id string, car *models.CarRequest) (*models.Car, error) {
	tracer := otel.Tracer("CarService")
	ctx, span := tracer.Start(ctx, "UpdateCar-Service")
	defer span.End()

	if err := models.ValidateRequest(*car); err != nil {
		return nil, err
	}

	updateCar, err := s.store.UpdateCar(ctx, id, car)
	if err != nil {
		return nil, err
	}

	return &updateCar, nil
}

func (s *CarService) DeleteCar(ctx context.Context, id string) (*models.Car, error) {
	tracer := otel.Tracer("CarService")
	ctx, span := tracer.Start(ctx, "DeleteCar-Service")
	defer span.End()

	deleteCar, err := s.store.DeleteCar(ctx, id)
	if err != nil {
		return nil, err
	}

	return &deleteCar, nil
}
