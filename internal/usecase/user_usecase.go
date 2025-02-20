package usecase

import "github.com/riskibarqy/ihsan-test/internal/domain"

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) GetUserByID(id int) (*domain.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *UserUsecase) CreateUser(user *domain.User) error {
	return uc.repo.Create(user)
}
