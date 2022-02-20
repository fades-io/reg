package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

type Server struct {
	service Service
	Router  *httprouter.Router
}

func (server *Server) Init(storage Storage) {
	server.service = NewService(storage)

	server.Router = httprouter.New()
	server.initRouters()
}

// Запускаем сервер, слушаем порт
func (server *Server) Run() {
	fmt.Println("Запуск сервера на хосте")
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	log.Fatal(http.ListenAndServe(host+":"+port, server.Router))
}
