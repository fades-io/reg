package middlewares

import (
	"net/http"

	"github.com/fades-io/reg/internal/apperror"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

// Middleware Добавляет заголовки запросу и обрабатывает ошибки.
func Middleware(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		apperror.HandleError(w, r, h)
	}
}
