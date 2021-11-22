package command

import "errors"

type RefreshCommand struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

var (
	ErrRefreshCommandInvalid = errors.New("refresh command is invalid")
)

func (cmd *RefreshCommand) Validate() error {
	if len(cmd.AccessToken) <= 0 || len(cmd.RefreshToken) <= 0 {
		return ErrRefreshCommandInvalid
	}

	return nil
}
