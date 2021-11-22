package mapper

import (
	"testing"

	cmd "gwi/assignment/feature/user/domain/command"

	"github.com/stretchr/testify/assert"
)

var mapper UserMapper = UserMapper{}

func TestToDataLayer(t *testing.T) {
	command := getCommand()

	response := mapper.ToDataLayerRegister(command)

	assert.Equal(t, command.Username, response.Username)
	assert.Equal(t, command.Password, response.Password)
	assert.Equal(t, command.Email, response.Email)
}

func getCommand() *cmd.RegisterCommand {
	return &cmd.RegisterCommand{
		Username: "username",
		Password: "password",
		Email:    "test@gmail.com",
	}
}
