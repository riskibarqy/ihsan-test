package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/riskibarqy/ihsan-test/datatransfers"
	"github.com/riskibarqy/ihsan-test/internal/constants"
	"github.com/riskibarqy/ihsan-test/internal/domain"
	utils "github.com/riskibarqy/ihsan-test/pkg"
)

type UserBalanceHistoryRepository struct {
	DB *sql.DB
}

func NewUserBalanceHistoryRepository(db *sql.DB) *UserBalanceHistoryRepository {
	return &UserBalanceHistoryRepository{DB: db}
}

// GetByID retrieves a userBalanceHistory by ID
func (r *UserBalanceHistoryRepository) GetByID(ctx context.Context, id int) (*domain.UserBalanceHistory, error) {
	userBalanceHistory := &domain.UserBalanceHistory{}
	err := r.DB.QueryRowContext(ctx, "SELECT id, user_id, no_rekening, previous_balance, new_balance, transaction_id, transaction_type, created_at FROM user_balance_histories WHERE id = $1", id).
		Scan(&userBalanceHistory.ID,
			&userBalanceHistory.UserID,
			&userBalanceHistory.NoRekening,
			&userBalanceHistory.PreviousBalance,
			&userBalanceHistory.NewBalance,
			&userBalanceHistory.TransactionID,
			&userBalanceHistory.TransactionType,
			&userBalanceHistory.CreatedAt,
		)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if userBalanceHistory not found
		}
		return nil, err
	}
	return userBalanceHistory, nil
}

// GetAll retrieves userBalanceHistories with cursor-based pagination
func (r *UserBalanceHistoryRepository) GetAll(ctx context.Context, params *datatransfers.ListQueryParams) ([]*domain.UserBalanceHistory, int, error) {
	query := "SELECT id, user_id, no_rekening, previous_balance, new_balance, transaction_id, transaction_type, created_at FROM user_balance_histories WHERE id > $1"
	args := []interface{}{params.Cursor}
	argIndex := 2

	query += fmt.Sprintf(" ORDER BY id ASC LIMIT $%d", argIndex)
	args = append(args, params.Limit)

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var userBalanceHistories []*domain.UserBalanceHistory
	var lastID int
	for rows.Next() {
		userBalanceHistory := &domain.UserBalanceHistory{}
		if err := rows.Scan(&userBalanceHistory.ID,
			&userBalanceHistory.UserID,
			&userBalanceHistory.NoRekening,
			&userBalanceHistory.PreviousBalance,
			&userBalanceHistory.NewBalance,
			&userBalanceHistory.TransactionID,
			&userBalanceHistory.TransactionType,
			&userBalanceHistory.CreatedAt,
		); err != nil {
			return nil, 0, err
		}
		userBalanceHistories = append(userBalanceHistories, userBalanceHistory)
		lastID = userBalanceHistory.ID
	}

	return userBalanceHistories, lastID, nil
}

// Create Insert data to user balance history
func (r *UserBalanceHistoryRepository) Create(userID int, amount int64, transactionID string, transactionType string) (int64, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	now := utils.Now()

	var updateQuery string
	switch transactionType {
	case string(constants.TransactionTypeAdd):
		updateQuery = `
			UPDATE public."users"
			SET "balance" = "balance" + $1
			WHERE "id" = $2
			RETURNING "balance", "no_rekening"
		`
	case string(constants.TransactionTypeWithdraw):
		updateQuery = `
			UPDATE public."users"
			SET "balance" = "balance" - $1
			WHERE "id" = $2
			RETURNING "balance", "no_rekening"
		`
	default:
		return 0, errors.New("invalid transaction type")
	}

	var newBalance, prevBalance int64
	var noRekening string
	if err = tx.QueryRow(updateQuery, amount, userID).Scan(&newBalance, &noRekening); err != nil {
		return 0, err
	}
	prevBalance = newBalance - amount

	insertHistoryQuery := `
		INSERT INTO public."user_balance_histories"
		("user_id", "no_rekening", "previous_balance", "change_amount", "new_balance", "transaction_type", "transaction_id", "created_at")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err = tx.Exec(insertHistoryQuery, userID, noRekening, prevBalance, amount, newBalance, transactionType, transactionID, now)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return newBalance, nil
}
