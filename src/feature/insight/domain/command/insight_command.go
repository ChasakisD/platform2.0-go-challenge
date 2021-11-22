package command

import (
	"errors"
)

type InsightCommand struct {
	Description string `json:"description"`
}

var (
	ErrInsightCommandInvalid = errors.New("insight is invalid")
)

func (cmd *InsightCommand) Validate() error {
	if len(cmd.Description) <= 0 {
		return ErrInsightCommandInvalid
	}

	return nil
}
