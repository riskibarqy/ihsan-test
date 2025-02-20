package usecase

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/riskibarqy/ihsan-test/datatransfers"
	"github.com/riskibarqy/ihsan-test/internal/constants"
	"github.com/riskibarqy/ihsan-test/internal/domain"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

// GetUserBalanceByNoRekening get user balance by no rekening
func (u *UserUsecase) GetUserBalanceByNoRekening(ctx context.Context, no_rekening string) (*datatransfers.UserGetBalanceResponse, error) {
	users, _, err := u.userRepo.GetAll(ctx, &datatransfers.ListQueryParams{
		NoRekening: no_rekening,
	})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New(constants.RecordNotFound)
	}

	if len(users) > 1 {
		log.Info("possible duplicate no rekening")
		return nil, errors.New(constants.UnknownError)
	}

	user := users[0]

	return &datatransfers.UserGetBalanceResponse{
		Saldo:      user.Balance,
		NoRekening: user.NoRekening,
	}, nil
}

// GetAll get all user by params
func (u *UserUsecase) GetAll(ctx context.Context, params *datatransfers.ListQueryParams) ([]*domain.User, int, error) {
	return u.userRepo.GetAll(ctx, params)
}

// CreateUser create user
func (u *UserUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	return u.userRepo.Create(ctx, user)
}
