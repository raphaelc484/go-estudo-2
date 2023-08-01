package config

import (
	"fmt"
	"os"
	"test/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		// Criar a pasta do banco de dados se não existir
		err := os.Mkdir("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		// Criar o arquivo do banco de dados se não existir
		file, err := os.Create(dbPath)
		if err != nil {
			fmt.Println("Erro ao criar o arquivo do banco de dados:", err)
			return nil, err
		}
		file.Close()
		fmt.Println("Arquivo do banco de dados criado com sucesso.")
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
