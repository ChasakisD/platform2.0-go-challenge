package mapper

import (
	ent "gwi/assignment/feature/user/data/entity"
	cmd "gwi/assignment/feature/user/domain/command"
)

type UserMapper struct{}

func (mapper *UserMapper) ToDataLayerRegister(user *cmd.RegisterCommand) *ent.User {
	return &ent.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password}
}
