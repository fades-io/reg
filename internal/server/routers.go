package server

import (
	"github.com/fades-io/reg/internal/middlewares"
	"net/http"
)

const (
	userLoginURL = "/reg"
)

func (server *Server) initRouters() {
	server.Router.HandlerFunc(http.MethodPost, userLoginURL, middlewares.SetHeadersMiddleware(server.Reg))
}
