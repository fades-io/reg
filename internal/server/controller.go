package server

import (
	"encoding/json"
	// TODO: присвоить
	"github.com/fades-io/reg/internal/apperror"
	"io/ioutil"
	"net/http"
)

// Это точно не нужно

// Обработка запрос для входа пользователя
func (server *Server) Login(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return apperror.New(err, "Не удалось считать тело запроса", err.Error(), http.StatusUnprocessableEntity)
	}

	user := domain.UserLogin{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return apperror.New(err, "Не удалось преобразовать JSON в модель", err.Error(), http.StatusUnprocessableEntity)
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		return apperror.New(err, "Неверный формат данных. Проверьте, корректно ли введен логин/пароль", err.Error(), http.StatusBadRequest)
	}
	token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		return apperror.SystemError(err)
	}
	return nil
}

func (server *Server) SignIn(username, password string) (string, error) {
	user, err := server.service.GetUser(username, password)
	if err != nil {
		return "", apperror.New(err, "Пользователя с таким логином/паролем не существует", err.Error(), http.StatusNotFound)
	}

	err = domain.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", apperror.New(err, "Неверный пароль", err.Error(), http.StatusUnauthorized)
	}

	return auth.CreateToken(user.ID)
}

// Регистрация нового пользователя
// TODO: implement
func (server *Server) Reg(w http.ResponseWriter, r *http.Request) error {
	// Получаем тело запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return apperror.New(err, "Не удалось считать тело запроса", err.Error(), http.StatusUnprocessableEntity)
	}

	user := domain.UserStruct{}
	// Преобразовываем JSON в модель
	err = json.Unmarshal(body, &user)
	if err != nil {
		return apperror.New(err, "Не удалось преобразовать JSON в модель", err.Error(), http.StatusUnprocessableEntity)
	}

	user.Prepare()
	err = user.Validate("reg")
	if err != nil {
		return apperror.New(err, "Неверный формат данных. Проверьте, корректно ли введен логин/пароль", err.Error(), http.StatusBadRequest)
	}
	// Вместо
	// token, err := server.SignIn(user.Username, user.Password)
	err = server.service.RegUser(user.Username, user.Email, user.Password)
	if err != nil {
		return "", apperror.New(err, "Пользователя с таким логином/паролем не существует", err.Error(), http.StatusNotFound)
	}
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	if err != nil {
		return apperror.SystemError(err)
	}
	return nil
}
