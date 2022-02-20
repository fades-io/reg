package postgresql

import (
	"github.com/fades-io/reg/internal/domain"
	"github.com/fades-io/reg/internal/server"
	"gorm.io/gorm"
)

// Обертка над gorm.DB
type postgresDB struct {
	db *gorm.DB
}

// Конструктор БД
func New(db *gorm.DB) server.Storage {
	return &postgresDB{
		db: db,
	}
}

// RegUser Регистрация нового пользователя
// TODO: implement
func (postgres *postgresDB) RegUser(username, email, password string) (*domain.User, error) {
	user := domain.User{}

	err := postgres.db.Debug().Create(&User{Username: username, Email: email, Password: password})

	if err != nil {
		return nil, err
	}

	return &user, nil
}
