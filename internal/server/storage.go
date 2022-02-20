package server

// Storage Интерфейс для работы с БД
type Storage interface {
	RegUser(username, email, password string) error
}
