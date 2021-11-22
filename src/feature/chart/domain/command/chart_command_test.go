package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChartCommandSuccess(t *testing.T) {
	command := ChartCommand{
		Description: "Asset Desc",
		XAxes:       "123",
		YAxes:       "1234",
		Points: []ChartPointCommand{
			{
				XValue: 1,
				YValue: 2,
			},
		},
	}

	assert.Nil(t, command.Validate())
}

func TestChartCommandEmptyDescription(t *testing.T) {
	command := ChartCommand{
		Description: "",
		XAxes:       "123",
		YAxes:       "1234",
		Points: []ChartPointCommand{
			{
				XValue: 1,
				YValue: 2,
			},
		},
	}

	assert.Equal(t, command.Validate(), ErrCreateChartCommandInvalid)
}

func TestChartCommandEmptyXAxes(t *testing.T) {
	command := ChartCommand{
		Description: "Asset Desc",
		XAxes:       "",
		YAxes:       "1234",
		Points: []ChartPointCommand{
			{
				XValue: 1,
				YValue: 2,
			},
		},
	}

	assert.Equal(t, command.Validate(), ErrCreateChartCommandInvalid)
}

func TestChartCommandEmptyYAxes(t *testing.T) {
	command := ChartCommand{
		Description: "Asset Desc",
		XAxes:       "123",
		YAxes:       "",
		Points: []ChartPointCommand{
			{
				XValue: 1,
				YValue: 2,
			},
		},
	}

	assert.Equal(t, command.Validate(), ErrCreateChartCommandInvalid)
}

func TestChartCommandEmptyPoints(t *testing.T) {
	command := ChartCommand{
		Description: "Asset Desc",
		XAxes:       "123",
		YAxes:       "1233",
	}

	assert.Equal(t, command.Validate(), ErrCreateChartCommandInvalid)
}
