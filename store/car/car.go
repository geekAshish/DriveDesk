package car

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/geekAshish/DriveDesk/models"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) GetCarById(ctx context.Context, id string) (models.Car, error) {
	car := models.Car{}

	query := `SELECT c.id, c.brand, c.model, c.year, c.color, c.price, c.is_engine c.created_at, c.updated_at e.id e.displacement e.no_of_cylinders e.fuel_type e.car_range FROM cars c LEFT JOIN engines e ON c.engine_id = e.id WHERE c.id = $1`

	row := s.db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&car.ID,
		&car.Name,
		&car.Year,
		&car.Brand,
		&car.FuelType,
		&car.Price,
		&car.CreateAt,
		&car.UpdateAt,

		&car.Engine.EngineID,
		&car.Engine.CarRange,
		&car.Engine.Dispacement,
		&car.Engine.NoOfCylinders,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return car, nil // No car found

		}
		return car, err
	}

	return car, nil
}

func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	var cars = []models.Car{}
	var query string

	if isEngine {
		query = `
		SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at, e.id, e.displacement, e.no_of_cylinders, e.car_range
		FROM cars c
		LEFT JOIN engines e ON c.engine_id = e.id
		WHERE c.brand = $1
		`
	} else {
		query = `
		SELECT id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at
		FROM car
		WHERE brand = $1
		`
	}

	rows, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var car models.Car

		if isEngine {
			var engine models.Engine
			err := rows.Scan(
				&car.ID,
				&car.Name,
				&car.Year,
				&car.Brand,
				&car.FuelType,
				&car.Engine.EngineID,
				&car.Price,
				&car.CreateAt,
				&car.UpdateAt,
				&car.Engine.Dispacement,
				&car.Engine.NoOfCylinders,
				&car.Engine.CarRange,
			)

			if err != nil {
				return nil, err
			}

			car.Engine = engine
		} else {
			err = rows.Scan(
				&car.ID,
				&car.Name,
				&car.Year,
				&car.Brand,
				&car.FuelType,
				&car.Engine.EngineID,
				&car.Price,
				&car.CreateAt,
				&car.UpdateAt,
			)

			if err != nil {
				return nil, err
			}
		}

		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

func (s Store) CreateCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error) {
	var createdCar models.Car
	var engineId uuid.UUID

	err := s.db.QueryRowContext(ctx, `SLECT id FROM engine WHERE id=$1`, carReq.Engine.EngineID).Scan(&engineId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return createdCar, errors.New("engine not found")
		}
		return createdCar, err
	}

	carID := uuid.New()

	create_at := time.Now()
	updated_at := time.Now()

	newCar := models.Car{
		ID:       carID,
		Name:     carReq.Name,
		Year:     carReq.Year,
		Brand:    carReq.Brand,
		FuelType: carReq.FuelType,
		Price:    carReq.Price,
		Engine:   carReq.Engine,
		CreateAt: create_at,
		UpdateAt: updated_at,
	}

	// begin the transaction , atomic [if we have any error in the middle, we will rollback]
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return createdCar, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}

		err = tx.Commit();
	}()

	query := `
	INSERT INTO cars (id, name, year, brand, fuel_type, price, engine_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURN id name year brand fuel_type price created_at updated_at
	`

	err = tx.QueryRowContext(ctx, query,
		newCar.ID,
		newCar.Name,
		newCar.Year,
		newCar.Brand,
		newCar.FuelType,
		newCar.Engine.EngineID,
		newCar.Price,
		newCar.CreateAt,
		newCar.UpdateAt,
	).Scan(
		&createdCar.ID,
		&createdCar.Name,
		&createdCar.Year,
		&createdCar.Brand,
		&createdCar.FuelType,
		&createdCar.Engine.EngineID,
		&createdCar.Price,
		&createdCar.CreateAt,
		&createdCar.UpdateAt,
	)

	if err != nil {
		return createdCar, err;
	}

	return createdCar, nil;
}

func (s Store) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (models.Car, error) {
	var updatedCar models.Car;

	// begin the transaction , atomic [if we have any error in the middle, we will rollback]
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return updatedCar, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}

		err = tx.Commit();
	}()

	query := `
	UPDATE cars
	SET name = $2, year = $3, brand = $4, fuel_type = $5, price = $6, engine_id = $7, updated_at = $8
	WHERE id = $1
	RETURN id, name, year, brand, fuel_type, price, created_at, updated_at
	`

	err = tx.QueryRowContext(ctx, query,
		id,
		carReq.Name,
		carReq.Year,
		carReq.Brand,
		carReq.FuelType,
		carReq.Engine.EngineID,
		carReq.Price,
		time.Now(),
	).Scan(
		&updatedCar.ID,
		&updatedCar.Name,
		&updatedCar.Year,
		&updatedCar.Brand, 
		&updatedCar.FuelType,
		&updatedCar.Engine.EngineID,
		&updatedCar.Price,
		&updatedCar.CreateAt,
		&updatedCar.UpdateAt,
	)

	if err != nil {
		return updatedCar, err;
	}

	return updatedCar, nil; 
}

func (s Store) DeleteCar(ctx context.Context, id string) (models.Car, error) {
	var deletedCar models.Car;

	// begin the transaction , atomic [if we have any error in the middle, we will rollback]
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return deletedCar, err;
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}

		err = tx.Commit();
	}()

	query := `
	SELECT id, name, year, brand, fuel_type, price, created_at, updated_at
	FROM car
	WHERE id = $1
	`;

	err = tx.QueryRowContext(ctx, query, id).Scan(
		&deletedCar.ID,
		&deletedCar.Name,
		&deletedCar.Year,
		&deletedCar.Brand,
		&deletedCar.FuelType,
		&deletedCar.Price,
		&deletedCar.CreateAt,
		&deletedCar.UpdateAt, 
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Car{}, errors.New("car not found")
		}
		return models.Car{}, err
	}

	result, err := tx.ExecContext(ctx, `DELETE FROM cars WHERE id = $1`, id)
	if err != nil {
		return models.Car{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Car{}, err
	}

	if rowsAffected == 0 {
		return models.Car{}, errors.New("no car deleted")
	} 

	return deletedCar, nil
}
