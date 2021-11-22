package command

import (
	"errors"
)

type ChartCommand struct {
	Description string              `json:"description"`
	XAxes       string              `json:"xAxes"`
	YAxes       string              `json:"yAxes"`
	Points      []ChartPointCommand `json:"points"`
}

type ChartPointCommand struct {
	XValue float64 `json:"xValue"`
	YValue float64 `json:"yValue"`
}

var (
	ErrCreateChartCommandInvalid = errors.New("chart is invalid")
)

func (cmd *ChartCommand) Validate() error {
	if len(cmd.Description) <= 0 || len(cmd.Points) <= 0 || len(cmd.XAxes) <= 0 || len(cmd.YAxes) <= 0 {
		return ErrCreateChartCommandInvalid
	}

	return nil
}
