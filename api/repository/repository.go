package repository

import (
	"context"
)

// Generic repo for our differents models
type Repository interface {
	Create(ctx context.Context, models interface{}) (string, error)
	Get(ctx context.Context, fields ...string) (interface{}, error)
	GetAll(ctx context.Context) ([]interface{}, error)
	Update(ctx context.Context, models interface{}) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}
