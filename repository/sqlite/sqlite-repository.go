package sqlite

import (
	"test/config"
	"test/model"

	"gorm.io/gorm"
)

type UserRepositorySQLite struct {
	db *gorm.DB
}

func NewUserRepositorySQLite() (*UserRepositorySQLite, error) {
	return &UserRepositorySQLite{
		db: config.GeSQLite(),
	}, nil
}

func (r *UserRepositorySQLite) Save(user model.User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepositorySQLite) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
