package api_rest

import (
	dbexec "github.com/hidratechh/GO_WMS/db/sqlc/exec"
	"github.com/labstack/echo/v4"
)

type Server struct {
	store  *dbexec.SQLStore
	router *echo.Echo
}

func (server *Server) routerMapping(router *echo.Echo) {
	router.POST("/users", server.createUser)
	router.POST("/whse", server.createWHSE)
}

func NewServer(store *dbexec.SQLStore) *Server {
	server := &Server{store: store}
	router := echo.New()

	server.routerMapping(router)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Start(address)
}
