package domain

type UserRepository interface {
	GetByID(id int) (*User, error)
	Create(user *User) error
}
