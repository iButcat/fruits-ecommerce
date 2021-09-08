package repository

import (
	"context"
)

// Generic repo for our differents models
type Repository interface {
	Create(ctx context.Context, models interface{}) (string, error)
	Get(ctx context.Context, models interface{}, fields map[string]interface{}) (interface{}, error)
	GetAll(ctx context.Context, models interface{}) (interface{}, error)
	Update(ctx context.Context, models interface{}, fields map[string]interface{}) (bool, error)
	Delete(ctx context.Context, models interface{}, id string) (bool, error)
}
