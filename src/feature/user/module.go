package user

import (
	"gwi/assignment/feature/user/api"
	"gwi/assignment/feature/user/data/mapper"
	"gwi/assignment/feature/user/data/repository"
	"gwi/assignment/feature/user/domain/usecase"

	"github.com/go-chi/chi/v5"
)

func Initialize(router *chi.Mux) {
	userMapper := &mapper.UserMapper{}
	userRepository := &repository.UserRepository{}
	userUseCase := usecase.Create(userMapper, userRepository)
	userController := api.UserController{UserUseCase: userUseCase}

	//Public Endpoints
	router.Group(func(router chi.Router) {
		router.Post("/auth/login", userController.Authenticate)
		router.Post("/auth/register", userController.CreateUser)
		router.Post("/auth/refresh", userController.Refresh)
	})
}
