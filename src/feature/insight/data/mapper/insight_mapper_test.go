package mapper

import (
	"testing"

	coreEnt "gwi/assignment/core/data/entity"
	assetEnt "gwi/assignment/feature/asset/data/entity"
	ent "gwi/assignment/feature/insight/data/entity"
	cmd "gwi/assignment/feature/insight/domain/command"

	"github.com/stretchr/testify/assert"
)

var mapper InsightMapper = InsightMapper{}

func TestToDomainLayer(t *testing.T) {
	insight := getInsight()

	response := mapper.ToDomainLayer(insight)

	assert.Equal(t, insight.Id, response.Id)
	assert.Equal(t, insight.AssetId, response.AssetId)
	assert.Equal(t, insight.Asset.Id, response.AssetId)
	assert.Equal(t, insight.Asset.Description, response.Description)
}

func TestToDomainLayerList(t *testing.T) {
	insights := []ent.Insight{
		*getInsight(),
	}

	responseInsights := mapper.ToDomainLayerList(&insights)

	assert.Equal(t, len(insights), len(*responseInsights))
}

func TestToDomainLayerPaging(t *testing.T) {
	insights := []ent.Insight{
		*getInsight(),
	}

	page := 1
	responseInsights := mapper.ToDomainLayerPaging(&insights, page)

	assert.Equal(t, page, responseInsights.Page)
	assert.Equal(t, len(insights), len(responseInsights.Results))
}

func TestToDataLayer(t *testing.T) {
	command := getCommand()

	response := mapper.ToDataLayer(command)

	assert.Equal(t, command.Description, response.Asset.Description)
	assert.NotEmpty(t, len(response.Id))
	assert.NotEmpty(t, len(response.AssetId))
	assert.NotEmpty(t, len(response.Asset.Id))
}

func getCommand() *cmd.InsightCommand {
	return &cmd.InsightCommand{
		Description: "Asset Desc",
	}
}

func getInsight() *ent.Insight {
	return &ent.Insight{
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
	}
}
