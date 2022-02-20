package server

// Тип для работы с БД
type service struct {
	storage Storage
}

// Service Сервис для работы с БД
type Service interface {
	RegUser(username, email, password string) error
}

// NewService Конструктор для создания сервиса
func NewService(storage Storage) Service {
	return &service{
		storage: storage,
	}
}

// RegUser Регистрация нового пользователя
func (s *service) RegUser(username, email, password string) error {
	return s.storage.RegUser(username, email, password)
}
