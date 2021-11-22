package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthCommandSuccess(t *testing.T) {
	command := AuthCommand{
		Email:    "test@test.com",
		Password: "123",
	}

	assert.Nil(t, command.Validate())
}

func TestAuthCommandMailEmpty(t *testing.T) {
	command := AuthCommand{
		Email:    "",
		Password: "123",
	}

	assert.Equal(t, command.Validate(), ErrAuthCommandInvalid)
}

func TestAuthCommandMailInvalid(t *testing.T) {
	command := AuthCommand{
		Email:    "com",
		Password: "123",
	}

	assert.Equal(t, command.Validate(), ErrAuthCommandInvalid)
}

func TestAuthCommandMailInvalid2(t *testing.T) {
	command := AuthCommand{
		Email:    "test.com",
		Password: "123",
	}

	assert.Equal(t, command.Validate(), ErrAuthCommandInvalid)
}

func TestAuthCommandMailInvalid3(t *testing.T) {
	command := AuthCommand{
		Email:    "@test.com",
		Password: "123",
	}

	assert.Equal(t, command.Validate(), ErrAuthCommandInvalid)
}

func TestAuthCommandPasswordEmpty(t *testing.T) {
	command := AuthCommand{
		Email:    "test@test.com",
		Password: "",
	}

	assert.Equal(t, command.Validate(), ErrAuthCommandInvalid)
}
