package engine

import (
	"context"
	"database/sql"

	"github.com/geekAshish/DriveDesk/models"
)

type EnginStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EnginStore {
	return &EnginStore{db: db}
}

func (s EnginStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {

}

func (s EnginStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {
	
}

func (s EnginStore) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {

}

func (s EnginStore) DeleteEngine(ctx context.Context, id string) (models.Car, error) {
	
}
 