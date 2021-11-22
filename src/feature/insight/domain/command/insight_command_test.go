package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsightCommandSuccess(t *testing.T) {
	command := InsightCommand{
		Description: "Asset Desc",
	}

	assert.Nil(t, command.Validate())
}

func TestInsightCommandEmptyDescription(t *testing.T) {
	command := InsightCommand{
		Description: "",
	}

	assert.Equal(t, command.Validate(), ErrInsightCommandInvalid)
}
