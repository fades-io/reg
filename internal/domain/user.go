package domain

import (
	"errors"
	"github.com/fades-io/reg/internal/apperror"
	"github.com/fades-io/reg/internal/logs"
	"html"
	"regexp"
	"strings"
)

// UserToDB Модель пользователя, которую поместим из БД
type UserToDB struct {
	ID       uint32
	Username string
	Email    string
	Password string
}

// UserStruct Данные пользователя для регистрации
type UserStruct struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Prepare Подготовка информации о новом пользователе
func (user *UserStruct) Prepare() {
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.Password = html.EscapeString(strings.TrimSpace(user.Password))
}

// Validate Валидация информации о пользователе
// TODO: добавить усиленную валидацию
func (user *UserStruct) Validate(action string) error {
	switch strings.ToLower(action) {
	case "reg":
		if user.Username == "" {
			return apperror.NewAppError(errors.New(strings.ToLower(logs.LoginRequired)), logs.LoginRequired, logs.LoginRequired, 422)
		}
		if user.Email == "" {
			return apperror.NewAppError(errors.New(strings.ToLower(logs.EmailRequired)), logs.EmailRequired, logs.EmailRequired, 422)
		}
		if user.Password == "" {
			return apperror.NewAppError(errors.New(strings.ToLower(logs.PasswordRequired)), logs.PasswordRequired, logs.PasswordRequired, 422)
		}
		return nil
	default:

		return nil
	}
}

var (
	ErrBadFormat = errors.New("invalid format")
	emailRegexp  = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func ValidateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}
