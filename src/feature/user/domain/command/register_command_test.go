package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterCommandSuccess(t *testing.T) {
	command := RegisterCommand{
		Username: "test",
		Email:    "test@test.com",
		Password: "12345678",
	}

	assert.Nil(t, command.Validate())
}

func TestRegisterCommandUsernameEmpty(t *testing.T) {
	command := RegisterCommand{
		Username: "",
		Email:    "test@test.com",
		Password: "12345678",
	}

	assert.Equal(t, command.Validate(), ErrRegisterCommandInvalid)
}

func TestRegisterCommandMailEmpty(t *testing.T) {
	command := RegisterCommand{
		Username: "test",
		Email:    "",
		Password: "12345678",
	}

	assert.Equal(t, command.Validate(), ErrRegisterCommandInvalid)
}

func TestRegisterCommandMailInvalid(t *testing.T) {
	command := RegisterCommand{
		Username: "test",
		Email:    "com",
		Password: "12345678",
	}

	assert.Equal(t, command.Validate(), ErrRegisterCommandInvalid)
}

func TestRegisterCommandMailInvalid2(t *testing.T) {
	command := RegisterCommand{
		Username: "test",
		Email:    "test.com",
		Password: "12345678",
	}

	assert.Equal(t, command.Validate(), ErrRegisterCommandInvalid)
}

func TestRegisterCommandMailInvalid3(t *testing.T) {
	command := RegisterCommand{
		Username: "test",
		Email:    "@test.com",
		Password: "12345678",
	}

	assert.Equal(t, command.Validate(), ErrRegisterCommandInvalid)
}

func TestRegisterCommandPasswordEmpty(t *testing.T) {
	command := RegisterCommand{
		Username: "test",
		Email:    "test@test.com",
		Password: "",
	}

	assert.Equal(t, command.Validate(), ErrRegisterPasswordInvalid)
}

func TestRegisterCommandPasswordBelow8(t *testing.T) {
	command := RegisterCommand{
		Username: "test",
		Email:    "test@test.com",
		Password: "1234567",
	}

	assert.Equal(t, command.Validate(), ErrRegisterPasswordInvalid)
}
