package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"

	"gwi/assignment/core/data/database"
	res "gwi/assignment/core/domain/response"
)

var (
	ErrParamNotValid      = errors.New("parameters are not valid")
	ErrSomethingWentWrong = errors.New("parameters are not valid")
)

func GetUserIdFromContext(ctx context.Context) (string, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return "", err
	}

	return claims["userId"].(string), nil
}

func RenderUnauthorized(r *http.Request) {
	render.Status(r, http.StatusUnauthorized)
}

func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res.ErrorResponse{Error: ErrSomethingWentWrong.Error()})
		return
	}

	if err.Error() == database.ErrSQLConnection.Error() {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res.ErrorResponse{Error: err.Error()})
		return
	}

	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, res.ErrorResponse{Error: err.Error()})
}

func RenderOk(w http.ResponseWriter, r *http.Request, response interface{}) {
	render.JSON(w, r, response)
}
