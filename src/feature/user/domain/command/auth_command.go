package command

import (
	"errors"
	"net/mail"
)

type AuthCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	ErrAuthCommandInvalid = errors.New("user command is invalid")
)

func (cmd *AuthCommand) Validate() error {
	if len(cmd.Email) <= 0 || len(cmd.Password) <= 0 {
		return ErrAuthCommandInvalid
	}

	if _, err := mail.ParseAddress(cmd.Email); err != nil {
		return ErrAuthCommandInvalid
	}

	return nil
}
