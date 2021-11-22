package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAudienceCommandSuccess(t *testing.T) {
	command := AudienceCommand{
		Description:   "Asset Desc",
		Gender:        "Male",
		BirthCountry:  "Greece",
		AgeGroupMin:   10,
		AgeGroupMax:   20,
		StatTypeValue: 2.0,
		StatType: AudienceStatTypeCommand{
			Title:          "123123",
			TitleFormatted: "123123123",
		},
	}

	assert.Nil(t, command.Validate())
}

func TestAudienceCommandEmptyTitle(t *testing.T) {
	command := AudienceCommand{
		Description:   "Asset Desc",
		Gender:        "Male",
		BirthCountry:  "Greece",
		AgeGroupMin:   10,
		AgeGroupMax:   20,
		StatTypeValue: 2.0,
		StatType: AudienceStatTypeCommand{
			Title:          "",
			TitleFormatted: "123123123",
		},
	}

	assert.Equal(t, command.Validate(), ErrAudienceCommandInvalid)
}

func TestAudienceCommandEmptyTitleFormatted(t *testing.T) {
	command := AudienceCommand{
		Description:   "Asset Desc",
		Gender:        "Male",
		BirthCountry:  "Greece",
		AgeGroupMin:   10,
		AgeGroupMax:   20,
		StatTypeValue: 2.0,
		StatType: AudienceStatTypeCommand{
			Title:          "123",
			TitleFormatted: "",
		},
	}

	assert.Equal(t, command.Validate(), ErrAudienceCommandInvalid)
}

func TestAudienceCommandEmptyDescription(t *testing.T) {
	command := AudienceCommand{
		Description:   "",
		Gender:        "Male",
		BirthCountry:  "Greece",
		AgeGroupMin:   10,
		AgeGroupMax:   20,
		StatTypeValue: 2.0,
		StatType: AudienceStatTypeCommand{
			Title:          "123",
			TitleFormatted: "123",
		},
	}

	assert.Equal(t, command.Validate(), ErrAudienceCommandInvalid)
}

func TestAudienceCommandInvalidGender(t *testing.T) {
	command := AudienceCommand{
		Description:   "",
		Gender:        "Male2",
		BirthCountry:  "Greece",
		AgeGroupMin:   10,
		AgeGroupMax:   20,
		StatTypeValue: 2.0,
		StatType: AudienceStatTypeCommand{
			Title:          "123",
			TitleFormatted: "123",
		},
	}

	assert.Equal(t, command.Validate(), ErrAudienceCommandInvalid)
}
