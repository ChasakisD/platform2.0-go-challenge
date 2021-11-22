package command

import (
	"errors"
)

type AudienceCommand struct {
	Description   string                  `json:"description"`
	Gender        string                  `json:"gender"`
	BirthCountry  string                  `json:"birthCountry"`
	AgeGroupMin   int                     `json:"ageGroupMin"`
	AgeGroupMax   int                     `json:"ageGroupMax"`
	StatType      AudienceStatTypeCommand `json:"statType"`
	StatTypeValue float64                 `json:"statTypeValue"`
}

type AudienceStatTypeCommand struct {
	Title          string `json:"title"`
	TitleFormatted string `json:"titleFormatted"`
}

var (
	ErrAudienceCommandInvalid = errors.New("audience is invalid")
)

func (cmd *AudienceCommand) Validate() error {
	if len(cmd.Description) <= 0 || len(cmd.StatType.Title) <= 0 || len(cmd.StatType.TitleFormatted) <= 0 {
		return ErrAudienceCommandInvalid
	}

	if len(cmd.Gender) > 0 && cmd.Gender != "Male" && cmd.Gender != "Female" {
		return ErrAudienceCommandInvalid
	}

	return nil
}
