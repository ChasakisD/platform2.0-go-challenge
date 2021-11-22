package domain

import (
	ent "gwi/assignment/feature/user/data/entity"
	cmd "gwi/assignment/feature/user/domain/command"
	res "gwi/assignment/feature/user/domain/response"
)

type UserRepository interface {
	CreateUser(user *ent.User) (*res.AuthResponse, error)
	Authenticate(email, password string) (*res.AuthResponse, error)
	Refresh(accessToken, refreshToken string) (*res.AuthResponse, error)
}

type UserMapper interface {
	ToDataLayerRegister(insight *cmd.RegisterCommand) *ent.User
}
