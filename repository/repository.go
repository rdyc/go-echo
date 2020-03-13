package repsitory

import (
	"context"

	"github.com/rdyc/go-echo/entities"
)

// UserRepo explain...
type UserRepo interface {
	Fetch(ctx context.Context, num int64) ([]*entities.User, error)
	GetByID(ctx context.Context, id int64) (*entities.User, error)
	Create(ctx context.Context, p *entities.User) (int64, error)
	Update(ctx context.Context, p *entities.User) (*entities.User, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
