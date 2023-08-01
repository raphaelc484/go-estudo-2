package usecase

import (
	"errors"
	"test/model"
	"test/repository"
)

type RegisterUserCase struct {
	userRepository repository.UserRepository
}

func NewRegisterUserCase(userRepository repository.UserRepository) *RegisterUserCase {
	return &RegisterUserCase{
		userRepository: userRepository,
	}
}

func (u *RegisterUserCase) RegisterUserCase(name, email string) (*model.User, error) {
	if name == "" {
		return nil, errors.New("nome não pode ser vazio")
	}

	if email == "" {
		return nil, errors.New("email não pode ser vazio")
	}

	user := model.User{
		Name:  name,
		Email: email,
	}

	err := u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
