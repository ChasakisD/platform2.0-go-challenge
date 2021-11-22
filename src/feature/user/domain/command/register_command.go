package command

import (
	"errors"
	"net/mail"
)

type RegisterCommand struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	ErrRegisterCommandInvalid  = errors.New("register command is invalid")
	ErrRegisterPasswordInvalid = errors.New("provide at least 8 characters for password")
)

func (cmd *RegisterCommand) Validate() error {
	if len(cmd.Username) <= 0 || len(cmd.Email) <= 0 {
		return ErrRegisterCommandInvalid
	}

	if _, err := mail.ParseAddress(cmd.Email); err != nil {
		return ErrRegisterCommandInvalid
	}

	if len(cmd.Password) < 8 {
		return ErrRegisterPasswordInvalid
	}

	return nil
}
