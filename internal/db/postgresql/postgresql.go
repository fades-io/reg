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

// NewDb Конструктор БД
func NewDb(db *gorm.DB) server.Storage {
	return &postgresDB{
		db: db,
	}
}

// RegUser Функция регистрации нового пользователя
func (postgres *postgresDB) RegUser(username, email, password string) error {

	err := postgres.db.Debug().Create(&domain.UserToDB{Username: username, Email: email, Password: password}).Error

	if err != nil {
		return err
	}

	return nil
}
