package api_rest

import (
	"net/http"

	db "github.com/hidratechh/GO_WMS/db/sqlc"
	"github.com/labstack/echo/v4"
)

type createWHSERequest struct {
	Whseid string `json:"whseid"`
	Desc   string `json:"desc"`
}

func (server *Server) createWHSE(ctx echo.Context) error {
	var req createWHSERequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	arg := db.CreateWHSEParams{
		Whseid: req.Whseid,
		Desc:   req.Desc,
	}
	whse, err := server.store.CreateWHSE(ctx.Request().Context(), arg)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)

	}
	return ctx.JSON(http.StatusCreated, whse)
}
