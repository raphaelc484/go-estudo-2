package inmemoryrepository

import (
	"test/model"
)

type UserRepositoryFake struct {
	users []*model.User
}

func NewUserRepositoryFake() *UserRepositoryFake {
	return &UserRepositoryFake{}
}

func (r *UserRepositoryFake) Save(user model.User) error {
	r.users = append(r.users, &user)
	return nil
}

func (r *UserRepositoryFake) FindByEmail(email string) (*model.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}
