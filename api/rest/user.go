package api_rest

import (
	"github.com/labstack/echo/v4"
)

type createUserRequest struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}

func (server *Server) createUser(ctx echo.Context) error {
	// var req createUserRequest
	// if err := ctx.Bind(&req); err != nil {
	// 	return ctx.JSON(http.StatusInternalServerError, err)
	// }

	// arg := db.CreateUserParams{
	// 	FullName: sql.NullString{String: req.FullName, Valid: true},
	// 	Username: sql.NullString{String: req.Username, Valid: true},
	// 	Password: sql.NullString{String: req.Password, Valid: true},
	// 	Photo:    sql.NullString{String: "/var/project/noqueue/img.jpg", Valid: true},
	// }
	// user, err := server.store.CreateUser(ctx.Request().Context(), arg)
	// if err != nil {
	// 	return ctx.JSON(http.StatusInternalServerError, err)

	// }
	// return ctx.JSON(http.StatusCreated, user)
	return nil
}
