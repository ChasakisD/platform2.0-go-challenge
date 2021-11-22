package mapper

import (
	"testing"

	coreEnt "gwi/assignment/core/data/entity"
	assetEnt "gwi/assignment/feature/asset/data/entity"
	ent "gwi/assignment/feature/chart/data/entity"
	cmd "gwi/assignment/feature/chart/domain/command"

	"github.com/stretchr/testify/assert"
)

var mapper ChartMapper = ChartMapper{}

func TestToDomainLayer(t *testing.T) {
	chart := getChart()

	response := mapper.ToDomainLayer(chart)

	assert.Equal(t, chart.Id, response.Id)
	assert.Equal(t, chart.AssetId, response.AssetId)
	assert.Equal(t, chart.Asset.Id, response.AssetId)
	assert.Equal(t, chart.Asset.Description, response.Description)
	assert.Equal(t, chart.XAxes, response.XAxes)
	assert.Equal(t, chart.YAxes, response.YAxes)
	assert.Equal(t, len(chart.Points), len(response.Points))
	assert.Equal(t, chart.Points[0].Id, response.Points[0].Id)
	assert.Equal(t, chart.Points[0].XValue, response.Points[0].XValue)
	assert.Equal(t, chart.Points[0].YValue, response.Points[0].YValue)
}

func TestToDomainLayerList(t *testing.T) {
	charts := []ent.Chart{
		*getChart(),
	}

	response := mapper.ToDomainLayerList(&charts)

	assert.Equal(t, len(charts), len(*response))
}

func TestToDomainLayerPaging(t *testing.T) {
	charts := []ent.Chart{
		*getChart(),
	}

	page := 1
	response := mapper.ToDomainLayerPaging(&charts, page)

	assert.Equal(t, page, response.Page)
	assert.Equal(t, len(charts), len(response.Results))
}

func TestToDataLayer(t *testing.T) {
	command := getCommand()

	response := mapper.ToDataLayer(command)

	assert.NotEmpty(t, len(response.Id))
	assert.NotEmpty(t, len(response.AssetId))
	assert.NotEmpty(t, len(response.Asset.Id))
	assert.NotEmpty(t, len(response.Points[0].Id))

	assert.Equal(t, command.Description, response.Asset.Description)
	assert.Equal(t, command.XAxes, response.XAxes)
	assert.Equal(t, command.YAxes, response.YAxes)
	assert.Equal(t, len(command.Points), len(response.Points))
	assert.Equal(t, command.Points[0].XValue, response.Points[0].XValue)
	assert.Equal(t, command.Points[0].YValue, response.Points[0].YValue)
}

func getCommand() *cmd.ChartCommand {
	return &cmd.ChartCommand{
		Description: "Asset Desc",
		XAxes:       "xAxes",
		YAxes:       "yAxes",
		Points: []cmd.ChartPointCommand{
			{
				XValue: 1.0,
				YValue: 5.0,
			},
		},
	}
}

func getChart() *ent.Chart {
	return &ent.Chart{
		BaseEntity: coreEnt.BaseEntity{
			Id: "123",
		},
		AssetId: "AssetId",
		Asset: assetEnt.Asset{
			BaseEntity: coreEnt.BaseEntity{
				Id: "AssetId",
			},
			Description: "Asset Desc",
		},
		XAxes: "xAxes",
		YAxes: "yAxes",
		Points: []ent.ChartPoint{
			{
				BaseEntity: coreEnt.BaseEntity{
					Id: "ChartId",
				},
				XValue: 1.0,
				YValue: 5.0,
			},
		},
	}
}
