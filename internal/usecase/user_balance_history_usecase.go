package usecase

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/riskibarqy/ihsan-test/datatransfers"
	"github.com/riskibarqy/ihsan-test/internal/constants"
	"github.com/riskibarqy/ihsan-test/internal/domain"
	utils "github.com/riskibarqy/ihsan-test/pkg"
)

type UserBalanceHistoryUsecase struct {
	userBalanceHistoryrepo domain.UserBalanceHistoryRepository
	userRepo               domain.UserRepository
}

func NewUserBalanceHistoryUsecase(userBalanceHistoryrepo domain.UserBalanceHistoryRepository, userRepo domain.UserRepository) *UserBalanceHistoryUsecase {
	return &UserBalanceHistoryUsecase{
		userBalanceHistoryrepo: userBalanceHistoryrepo,
		userRepo:               userRepo,
	}
}

// CreateBalance to add/withdraw funds to user accounts
func (u *UserBalanceHistoryUsecase) CreateBalance(ctx context.Context, params *datatransfers.UserBalanceHistoryTXRequest) (*datatransfers.UserBalanceHistoryTXResponse, error) {
	users, _, err := u.userRepo.GetAll(ctx, &datatransfers.ListQueryParams{
		NoRekening: params.NoRekening,
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

	if params.TransactionType == constants.TransactionTypeWithdraw {
		if user.Balance < params.Nominal {
			err := errors.New("insufficient balance")
			return nil, err
		}
	}

	newBalance, err := u.userBalanceHistoryrepo.Create(user.ID, int64(params.Nominal), utils.GenerateTransactionID(), string(params.TransactionType))
	if err != nil {
		return nil, err
	}

	return &datatransfers.UserBalanceHistoryTXResponse{
		Saldo:      int(newBalance),
		NoRekening: user.NoRekening,
	}, nil
}
