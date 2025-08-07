package engine

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/geekAshish/DriveDesk/models"
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

		err = tx.Commit();
	}()

	return engine, nil;
}

func (e EnginStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {
	
}

func (e EnginStore) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {

}

func (e EnginStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {
	
}
