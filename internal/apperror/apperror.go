package apperror

import (
	"encoding/json"
	"github.com/fades-io/reg/internal/logs"
)

//Готовые переменные для стандратных ошибок

var (
	ErrNotFound = NewAppError(nil, logs.NotFound, "", 404)
)

// AppError Кастомная ошибка, которая передается в json вместе с сообщением
// Специальная тип для ошибок в приложении
type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             uint32 `json:"code,omitempty"`
}

// NewAppError Функция создания новой ошибки
func NewAppError(err error, message, developerMessage string, code uint32) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

// Error метод для соответствия стандартному интерфейсу Error{}
func (appError *AppError) Error() string {
	return appError.Message
}

// Unwrap Возвращает корневую ошибку
func (appError *AppError) Unwrap() error {
	return appError.Err
}

// Marshal Вспомогательный метод для сериализации ошибки
func (appError *AppError) Marshal() []byte {
	marshal, err := json.Marshal(appError)
	if err != nil {
		return nil
	}
	return marshal
}

// SystemError Системная ошибка
func SystemError(err error) *AppError {
	return NewAppError(err, logs.InternalServerError, err.Error(), 418)
}
