package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/riskibarqy/ihsan-test/datatransfers"
	"github.com/riskibarqy/ihsan-test/internal/domain"
	utils "github.com/riskibarqy/ihsan-test/pkg"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	user := &domain.User{}
	err := r.DB.QueryRow("SELECT id, nama, nik, no_hp, no_rekening FROM users WHERE id = $1 and deleted_at is null", id).
		Scan(&user.ID, &user.Nama, &user.Nik, &user.NoHp, &user.NoRekening)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if user not found
		}
		return nil, err
	}
	return user, nil
}

// GetAll retrieves users with cursor-based pagination
func (r *UserRepository) GetAll(ctx context.Context, params *datatransfers.ListQueryParams) ([]*domain.User, int, error) {
	query := "SELECT id, nama, nik, no_hp, no_rekening FROM users WHERE id > $1"
	args := []interface{}{params.Cursor}
	argIndex := 2

	if params.Nama != "" {
		query += fmt.Sprintf(" AND nama ILIKE $%d", argIndex)
		args = append(args, "%"+params.Nama+"%")
		argIndex++
	}

	if params.NoRekening != "" {
		query += fmt.Sprintf(" AND no_rekening = $%d", argIndex)
		args = append(args, params.NoRekening)
		argIndex++
	}

	if params.NoHP != "" {
		query += fmt.Sprintf(" AND no_hp = $%d", argIndex)
		args = append(args, params.NoHP)
		argIndex++
	}

	if params.NIK != "" {
		query += fmt.Sprintf(" AND nik = $%d", argIndex)
		args = append(args, params.NIK)
		argIndex++
	}

	query += fmt.Sprintf(" AND deleted_at IS NULL ORDER BY id ASC LIMIT $%d", argIndex)
	args = append(args, params.Limit)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []*domain.User
	var lastID int
	for rows.Next() {
		user := &domain.User{}
		if err := rows.Scan(&user.ID, &user.Nama, &user.Nik, &user.NoHp, &user.NoRekening); err != nil {
			return nil, 0, err
		}
		users = append(users, user)
		lastID = user.ID
	}

	return users, lastID, nil
}

// Create inserts a new user
func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	now := utils.Now()
	_, err := r.DB.Exec(
		"INSERT INTO users (nama, nik, no_hp, no_rekening, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
		user.Nama, user.Nik, user.NoHp, user.NoRekening, now, now,
	)
	return err
}

// Update modifies an existing user
func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	now := utils.Now()
	result, err := r.DB.Exec(
		"UPDATE users SET nama = $1, nik = $2, no_hp = $3, no_rekening = $4, updated_at = $5 WHERE id = $6 and deleted_at is null",
		user.Nama, user.Nik, user.NoHp, user.NoRekening, now, user.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows updated")
	}

	return nil
}

// Delete removes a user by ID
func (r *UserRepository) Delete(ctx context.Context, id int) error {
	result, err := r.DB.Exec("UPDATE users set deleted_at = $1 where id = %2", utils.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no user deleted")
	}

	return nil
}
