package api

import (
	"encoding/json"
	"net/http"

	apiCore "gwi/assignment/core/api"
	coreRes "gwi/assignment/core/domain/response"
	cmd "gwi/assignment/feature/user/domain/command"
	res "gwi/assignment/feature/user/domain/response"
	useCase "gwi/assignment/feature/user/domain/usecase"
)

type UserController struct {
	UserUseCase *useCase.UserUseCase
}

func dummySwaggerModels() {
	print(coreRes.ErrorResponse{}.Error)
	print(res.AuthResponse{}.AccessToken)
}

// @Summary Create User
// @Description Creates a user
// @Tags User
// @Accept json
// @Produce json
// @Param command body cmd.RegisterCommand false "Command"
// @Success 200 {object} res.AuthResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 500 {object} coreRes.ErrorResponse
// @Router /auth/register [post]
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	command := &cmd.RegisterCommand{}
	err := json.NewDecoder(r.Body).Decode(command)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err = command.Validate(); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	authResponse, err := uc.UserUseCase.CreateUser(command)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, authResponse)
}

// @Summary Authenticate
// @Description Authenticate
// @Tags User
// @Accept json
// @Produce json
// @Param command body cmd.AuthCommand false "Command"
// @Success 200 {object} res.AuthResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Router /auth/login [post]
func (uc *UserController) Authenticate(w http.ResponseWriter, r *http.Request) {
	command := &cmd.AuthCommand{}
	err := json.NewDecoder(r.Body).Decode(command)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err = command.Validate(); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	authResponse, err := uc.UserUseCase.Authenticate(command.Email, command.Password)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, authResponse)
}

// @Summary Authenticate
// @Description Authenticate
// @Tags User
// @Accept json
// @Produce json
// @Param command body cmd.RefreshCommand false "Command"
// @Success 200 {object} res.AuthResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Router /auth/refresh [post]
func (uc *UserController) Refresh(w http.ResponseWriter, r *http.Request) {
	command := &cmd.RefreshCommand{}
	err := json.NewDecoder(r.Body).Decode(command)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err = command.Validate(); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	authResponse, err := uc.UserUseCase.Refresh(command.AccessToken, command.RefreshToken)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, authResponse)
}
