package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefreshCommandSuccess(t *testing.T) {
	command := RefreshCommand{
		AccessToken:  "test",
		RefreshToken: "test",
	}

	assert.Nil(t, command.Validate())
}

func TestRefreshCommandAccessEmpty(t *testing.T) {
	command := RefreshCommand{
		AccessToken:  "",
		RefreshToken: "test",
	}

	assert.Equal(t, command.Validate(), ErrRefreshCommandInvalid)
}

func TestRefreshCommandRefreshEmpty(t *testing.T) {
	command := RefreshCommand{
		AccessToken:  "test",
		RefreshToken: "",
	}

	assert.Equal(t, command.Validate(), ErrRefreshCommandInvalid)
}
