package domain

import (
	"context"

	"github.com/riskibarqy/ihsan-test/datatransfers"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*User, error)
	GetAll(ctx context.Context, params *datatransfers.ListQueryParams) ([]*User, int, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}
