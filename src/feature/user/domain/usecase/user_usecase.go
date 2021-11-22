package usecase

import (
	domain "gwi/assignment/feature/user/domain"
	cmd "gwi/assignment/feature/user/domain/command"
	res "gwi/assignment/feature/user/domain/response"
)

type UserUseCase struct {
	userMapper     domain.UserMapper
	userRepository domain.UserRepository
}

func Create(userMapper domain.UserMapper, userRepository domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		userMapper:     userMapper,
		userRepository: userRepository,
	}
}

func (useCase *UserUseCase) CreateUser(cmd *cmd.RegisterCommand) (*res.AuthResponse, error) {
	return useCase.userRepository.CreateUser(useCase.userMapper.ToDataLayerRegister(cmd))
}

func (useCase *UserUseCase) Authenticate(email, password string) (*res.AuthResponse, error) {
	return useCase.userRepository.Authenticate(email, password)
}

func (useCase *UserUseCase) Refresh(accessToken, refreshToken string) (*res.AuthResponse, error) {
	return useCase.userRepository.Refresh(accessToken, refreshToken)
}
