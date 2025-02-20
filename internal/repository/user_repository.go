package repository

import (
	"database/sql"

	"github.com/riskibarqy/ihsan-test/internal/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByID(id int) (*domain.User, error) {
	user := &domain.User{}
	err := r.DB.QueryRow("SELECT id, nama, nik, no_hp FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Nama, &user.Nik, &user.NoHp)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Create(user *domain.User) error {
	_, err := r.DB.Exec("INSERT INTO users (nama, nik, no_hp) VALUES ($1, $2)", user.Nama, user.Nik, user.NoHp)
	return err
}
