package server

import (
	"encoding/json"
	"github.com/fades-io/reg/internal/domain"

	"github.com/fades-io/reg/internal/apperror"
	"io/ioutil"
	"net/http"
)

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
	// token, apperror := server.SignIn(user.Username, user.Password)
	err = server.service.RegUser(user.Username, user.Email, user.Password)
	if err != nil {
		return apperror.New(err, "Неверный формат данных. Проверьте, корректно ли введен логин/пароль", err.Error(), http.StatusBadRequest)
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
