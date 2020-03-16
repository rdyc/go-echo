package repsitory

import (
	"context"

	"github.com/google/uuid"
	"github.com/rdyc/go-echo/entities"
)

// UserRepo explain...
type UserRepo interface {
	Fetch(ctx context.Context, num int64) ([]*entities.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	Create(ctx context.Context, p *entities.User) (uuid.UUID, error)
	Update(ctx context.Context, p *entities.User) (*entities.User, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, error)
}
