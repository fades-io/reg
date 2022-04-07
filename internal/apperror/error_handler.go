package apperror

import (
	"errors"
	"github.com/fades-io/reg/internal/middlewares"
	"net/http"
)

// HandleError Обрабатывает ошибку, которую вернул handler
func HandleError(w http.ResponseWriter, r *http.Request, h middlewares.AppHandler) error {
	var appErr AppError
	err := h(w, r)
	if err != nil {
		// Если специальная ошибка
		if errors.As(err, &appErr) {
			// Если ошибка - "не найдено"
			if errors.Is(err, ErrNotFound) {
				w.WriteHeader(http.StatusNotFound)
				_, err := w.Write(ErrNotFound.Marshal())
				if err != nil {
					return err
				} else {
					return nil
				}

			}
			// В случае других ошибок - подозреваем проблему с пользовательскими данными
			err = err.(*AppError)
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write(appErr.Marshal())
			if err != nil {
				return err
			} else {
				return nil
			}
		}

		// Оборачиваем все остальные ошибки в системные
		w.WriteHeader(http.StatusTeapot)
		_, err := w.Write(SystemError(err).Marshal())
		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return nil
	}
}
