package domain

import (
	"context"

	"github.com/riskibarqy/ihsan-test/datatransfers"
)

type UserBalanceHistoryRepository interface {
	GetByID(ctx context.Context, id int) (*UserBalanceHistory, error)
	GetAll(ctx context.Context, params *datatransfers.ListQueryParams) ([]*UserBalanceHistory, int, error)
	Create(userID int, amount int64, transactionID string, transactionType string) (int64, error)
}
