package middlewares

import (
	"github.com/fades-io/reg/internal/logs"
	"log"
	"net/http"

	"github.com/fades-io/reg/internal/apperror"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

// Middleware Добавляет заголовки запросу и обрабатывает ошибки.
func Middleware(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := apperror.HandleError(w, r, h)
		if err != nil {
			log.Printf(logs.FailedHandleError, err)
		}
	}
}
