package mapper

import (
	"testing"

	coreEnt "gwi/assignment/core/data/entity"
	assetEnt "gwi/assignment/feature/asset/data/entity"
	ent "gwi/assignment/feature/audience/data/entity"
	cmd "gwi/assignment/feature/audience/domain/command"

	"github.com/stretchr/testify/assert"
)

var mapper AudienceMapper = AudienceMapper{}

func TestToDomainLayer(t *testing.T) {
	audience := getAudience()

	response := mapper.ToDomainLayer(audience)

	assert.Equal(t, audience.Id, response.Id)
	assert.Equal(t, audience.AssetId, response.AssetId)
	assert.Equal(t, audience.Asset.Id, response.AssetId)
	assert.Equal(t, audience.Asset.Description, response.Description)
	assert.Equal(t, audience.Gender, response.Gender)
	assert.Equal(t, audience.BirthCountry, response.BirthCountry)
	assert.Equal(t, audience.AgeGroupMin, response.AgeGroupMin)
	assert.Equal(t, audience.AgeGroupMax, response.AgeGroupMax)
	assert.Equal(t, audience.StatTypeId, response.StatType.Id)
	assert.Equal(t, audience.StatType.Id, response.StatType.Id)
	assert.Equal(t, audience.StatTypeValue, response.StatTypeValue)
}

func TestToDomainLayerList(t *testing.T) {
	audiences := []ent.Audience{
		*getAudience(),
	}

	response := mapper.ToDomainLayerList(&audiences)

	assert.Equal(t, len(audiences), len(*response))
}

func TestToDomainLayerPaging(t *testing.T) {
	audiences := []ent.Audience{
		*getAudience(),
	}

	page := 1
	response := mapper.ToDomainLayerPaging(&audiences, page)

	assert.Equal(t, page, response.Page)
	assert.Equal(t, len(audiences), len(response.Results))
}

func TestToDataLayer(t *testing.T) {
	command := getCommand()

	response := mapper.ToDataLayer(command)

	assert.NotEmpty(t, len(response.Id))
	assert.NotEmpty(t, len(response.AssetId))
	assert.NotEmpty(t, len(response.Asset.Id))
	assert.NotEmpty(t, len(response.StatTypeId))
	assert.NotEmpty(t, len(response.StatType.Id))

	assert.Equal(t, command.Description, response.Asset.Description)
	assert.Equal(t, command.Gender, response.Gender)
	assert.Equal(t, command.BirthCountry, response.BirthCountry)
	assert.Equal(t, command.AgeGroupMin, response.AgeGroupMin)
	assert.Equal(t, command.AgeGroupMax, response.AgeGroupMax)
	assert.Equal(t, command.StatTypeValue, response.StatTypeValue)
}

func getCommand() *cmd.AudienceCommand {
	return &cmd.AudienceCommand{
		Description:   "Asset Desc",
		Gender:        "Male",
		BirthCountry:  "Greece",
		AgeGroupMin:   10,
		AgeGroupMax:   20,
		StatTypeValue: 2.0,
		StatType: cmd.AudienceStatTypeCommand{
			Title:          "123123",
			TitleFormatted: "123123123",
		},
	}
}

func getAudience() *ent.Audience {
	return &ent.Audience{
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
		Gender:        "Male",
		BirthCountry:  "Greece",
		AgeGroupMin:   10,
		AgeGroupMax:   20,
		StatTypeValue: 2.0,
		StatTypeId:    "StatTypeId",
		StatType: ent.AudienceStatType{
			BaseEntity: coreEnt.BaseEntity{
				Id: "StatTypeId",
			},
			Title:          "123123",
			TitleFormatted: "123123123",
		},
	}
}
