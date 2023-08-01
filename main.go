package main

import (
	"fmt"
	"test/config"
	"test/repository/sqlite"
	usecase "test/use-case"
)

func main() {
	config.Init()

	userRepositorySQLite, err := sqlite.NewUserRepositorySQLite()
	if err != nil {
		fmt.Println("Erro ao criar o repositório do SQLite:", err)
		return
	}

	// Criando o use case com o repositório do SQLite
	registerUseCase := usecase.NewRegisterUserCase(userRepositorySQLite)

	newUser, err := registerUseCase.RegisterUserCase("Fuji", "fuji@test.com")

	if err != nil {
		fmt.Println("Erro ao registrar usuário: ", err)
	} else {
		fmt.Printf("Usuário registrado com sucesso: %s (Email: %s)\n", newUser.Name, newUser.Email)
	}
}
