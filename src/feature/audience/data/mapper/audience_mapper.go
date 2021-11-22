package mapper

import (
	assetEnt "gwi/assignment/feature/asset/data/entity"
	ent "gwi/assignment/feature/audience/data/entity"
	cmd "gwi/assignment/feature/audience/domain/command"
	res "gwi/assignment/feature/audience/domain/response"

	"github.com/google/uuid"
)

type AudienceMapper struct{}

func (mapper *AudienceMapper) ToDomainLayer(audience *ent.Audience) *res.AudienceResponse {
	return &res.AudienceResponse{
		Id:           audience.Id,
		AssetId:      audience.AssetId,
		Gender:       audience.Gender,
		BirthCountry: audience.BirthCountry,
		AgeGroupMin:  audience.AgeGroupMin,
		AgeGroupMax:  audience.AgeGroupMax,
		StatType: res.AudienceStatTypeResponse{
			Id:             audience.StatType.Id,
			Title:          audience.StatType.Title,
			TitleFormatted: audience.StatType.TitleFormatted,
		},
		StatTypeValue:        audience.StatTypeValue,
		Description:          audience.Asset.Description,
		DescriptionFormatted: audience.GenerateDescription(),
	}
}

func (mapper *AudienceMapper) ToDomainLayerList(audiences *[]ent.Audience) *[]res.AudienceResponse {
	ret := []res.AudienceResponse{}
	for _, audience := range *audiences {
		ret = append(ret, *mapper.ToDomainLayer(&audience))
	}

	return &ret
}

func (mapper *AudienceMapper) ToDomainLayerPaging(audiences *[]ent.Audience, page int) *res.AudiencePageResponse {
	return &res.AudiencePageResponse{
		Page:    page,
		Results: *mapper.ToDomainLayerList(audiences),
	}
}

func (mapper *AudienceMapper) ToDataLayer(audience *cmd.AudienceCommand) *ent.Audience {
	assetEntity := &assetEnt.Asset{
		Description: audience.Description,
	}
	assetEntity.Id = uuid.NewString()

	statType := &ent.AudienceStatType{
		Title:          audience.StatType.Title,
		TitleFormatted: audience.StatType.TitleFormatted,
	}
	statType.Id = uuid.NewString()

	audienceEntity := &ent.Audience{
		Gender:        audience.Gender,
		BirthCountry:  audience.BirthCountry,
		AgeGroupMin:   audience.AgeGroupMin,
		AgeGroupMax:   audience.AgeGroupMax,
		Asset:         *assetEntity,
		AssetId:       assetEntity.Id,
		StatType:      *statType,
		StatTypeId:    statType.Id,
		StatTypeValue: audience.StatTypeValue,
	}
	audienceEntity.Id = uuid.NewString()

	return audienceEntity
}
