package mapper

import (
	assetEnt "gwi/assignment/feature/asset/data/entity"
	ent "gwi/assignment/feature/insight/data/entity"
	cmd "gwi/assignment/feature/insight/domain/command"
	res "gwi/assignment/feature/insight/domain/response"

	"github.com/google/uuid"
)

type InsightMapper struct{}

func (mapper *InsightMapper) ToDomainLayer(insight *ent.Insight) *res.InsightResponse {
	return &res.InsightResponse{
		Id:          insight.Id,
		AssetId:     insight.AssetId,
		Description: insight.Asset.Description,
	}
}

func (mapper *InsightMapper) ToDomainLayerList(insights *[]ent.Insight) *[]res.InsightResponse {
	ret := []res.InsightResponse{}
	for _, audience := range *insights {
		ret = append(ret, *mapper.ToDomainLayer(&audience))
	}

	return &ret
}

func (mapper *InsightMapper) ToDomainLayerPaging(insights *[]ent.Insight, page int) *res.InsightPageResponse {
	return &res.InsightPageResponse{
		Page:    page,
		Results: *mapper.ToDomainLayerList(insights),
	}
}

func (mapper *InsightMapper) ToDataLayer(audience *cmd.InsightCommand) *ent.Insight {
	assetEntity := &assetEnt.Asset{
		Description: audience.Description,
	}
	assetEntity.Id = uuid.NewString()

	insightEntity := &ent.Insight{
		Asset:   *assetEntity,
		AssetId: assetEntity.Id,
	}
	insightEntity.Id = uuid.NewString()

	return insightEntity
}
