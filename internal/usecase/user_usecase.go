package usecase

import (
	"context"

	"github.com/riskibarqy/ihsan-test/datatransfers"
	"github.com/riskibarqy/ihsan-test/internal/domain"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *UserUsecase) GetAll(ctx context.Context, params *datatransfers.ListQueryParams) ([]*domain.User, int, error) {
	return u.repo.GetAll(ctx, params)
}

func (u *UserUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	return u.repo.Create(ctx, user)
}
