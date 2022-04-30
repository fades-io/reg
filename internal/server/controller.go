package server

import (
	"encoding/json"
	"github.com/fades-io/reg/internal/apperror"
	"github.com/fades-io/reg/internal/domain"
	"io/ioutil"
	"net/http"
)

// Reg Регистрация нового пользователя
func (server *Server) Reg(w http.ResponseWriter, r *http.Request) error {
	// Получаем тело запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return apperror.NewAppError(err, apperror.InvalidRequestBody, err.Error(), http.StatusUnprocessableEntity)
	}

	user := domain.UserStruct{}
	// Преобразовываем JSON в ошибку
	err = json.Unmarshal(body, &user)
	if err != nil {
		return apperror.NewAppError(err, apperror.JsonParsingError, err.Error(), http.StatusUnprocessableEntity)
	}

	user.Prepare()
	err = user.Validate("reg")
	if err != nil {
		return apperror.NewAppError(err, apperror.InvalidFormat, err.Error(), http.StatusBadRequest)
	}

	err = server.service.RegUser(user.Username, user.Email, user.Password)
	if err != nil {
		return apperror.NewAppError(err, apperror.InvalidFormat, err.Error(), http.StatusBadRequest)
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
