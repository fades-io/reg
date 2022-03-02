package server

import (
	"github.com/fades-io/reg/internal/middlewares"
	"net/http"
)

const (
	userRegURL = "/reg"
)

func (server *Server) initRouters() {
	server.Router.HandlerFunc(http.MethodPost, userRegURL, middlewares.SetHeadersMiddleware(server.Reg))
}
