package repository

import "test/model"

type UserRepository interface {
	Save(user model.User) error
	FindByEmail(email string) (*model.User, error)
}
