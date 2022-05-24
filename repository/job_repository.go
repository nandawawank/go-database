package repository

import (
	"context"
	"godatabase/entity"
)

type JobRepository interface {
	Insert(ctx context.Context, job entity.Job) (entity.Job, error)
	FindById(ctx context.Context, id int32) (entity.Job, error)
	FindAll(ctx context.Context) ([]entity.Job, error)
}
