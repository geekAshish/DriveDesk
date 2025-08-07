package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/geekAshish/DriveDesk/models"
	"github.com/google/uuid"
)

type EnginStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EnginStore {
	return &EnginStore{db: db}
}

func (e EnginStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine

	// begin the transaction , atomic [if we have any error in the middle, we will rollback]
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("transaction rollback error: %v/n", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("transaction commit error: %v/n", cmErr)
			}
		}

		err = tx.Commit()
	}()

	err = tx.QueryRowContext(ctx, `SELECT id, displacement, no_of_cyclinder, car_range FROM engine WHERE id=$1`, id).Scan(
		&engine.EngineID,
		&engine.Dispacement,
		&engine.NoOfCylinders,
		&engine.CarRange,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return engine, errors.New("engine with id not found")
		}
		return engine, err
	}

	return engine, nil
}

func (e EnginStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("transaction rollback error: %v/n", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("transaction commit error: %v/n", cmErr)
			}
		}

		err = tx.Commit()
	}()

	engineID := uuid.New()

	_, err = tx.ExecContext(
		ctx,
		`INSERT INTO engine (id, displacement, no_of_cyclinder, car_range) VALUES ($1, $2, $3, $4)`,
		engineID, engineReq.Dispacement, engineReq.NoOfCylinders, engineReq.CarRange,
	)

	if err != nil {
		return models.Engine{}, nil
	}

	engine := models.Engine{
		EngineID:      engineID,
		Dispacement:   engineReq.Dispacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}

	return engine, nil
}

func (e EnginStore) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {
	enginID, err := uuid.Parse(id);
	if err != nil {
		return models.Engine{}, fmt.Errorf("invalid engine id format: %w", id)
	}

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("transaction rollback error: %v/n", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("transaction commit error: %v/n", cmErr)
			}
		}

		err = tx.Commit()
	}()

	result, err := tx.ExecContext(
		ctx,
		`UPDATE engine SET displacement=$1, no_of_cyclinder=$2, car_range=$3 WHERE id=$4`,
		engineReq.Dispacement,
		engineReq.NoOfCylinders,
		engineReq.CarRange,
	)

	if err != nil {
		return models.Engine{}, err;
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return models.Engine{}, err
	}
	if rowAffected == 0 {
		return models.Engine{}, errors.New("engine with id not found")
	}

	engine := models.Engine{
		EngineID:      enginID, 
		Dispacement:   engineReq.Dispacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}

	return engine, nil
}

func (e EnginStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {

}
