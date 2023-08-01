package usecase_test

import (
	inmemoryrepository "test/repository/in-memory-repository"
	usecase "test/use-case"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUserCase_Register(t *testing.T) {
	fake := &inmemoryrepository.UserRepositoryFake{}
	useCase := usecase.NewRegisterUserCase(fake)

	name := "Joe"
	email := "joe@test.com"

	newUser, err := useCase.RegisterUserCase(name, email)

	assert.Nil(t, err)
	assert.NotNil(t, newUser)
	assert.Equal(t, name, newUser.Name)
	assert.Equal(t, email, newUser.Email)
}

func TestRegisterUserCase_RegisterInvalidName(t *testing.T) {
	fake := &inmemoryrepository.UserRepositoryFake{}
	useCase := usecase.NewRegisterUserCase(fake)

	name := ""
	email := "joe@test.com"

	newUser, err := useCase.RegisterUserCase(name, email)

	assert.NotNil(t, err)
	assert.Nil(t, newUser)
	assert.Equal(t, "nome não pode ser vazio", err.Error())
}

func TestRegisterUserCase_RegisterInvalidEmail(t *testing.T) {
	fake := &inmemoryrepository.UserRepositoryFake{}
	useCase := usecase.NewRegisterUserCase(fake)

	name := "Joe"
	email := ""

	newUser, err := useCase.RegisterUserCase(name, email)

	assert.NotNil(t, err)
	assert.Nil(t, newUser)
	assert.Equal(t, "email não pode ser vazio", err.Error())
}
