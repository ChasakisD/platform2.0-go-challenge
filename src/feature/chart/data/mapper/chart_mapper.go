package mapper

import (
	assetEnt "gwi/assignment/feature/asset/data/entity"
	ent "gwi/assignment/feature/chart/data/entity"
	cmd "gwi/assignment/feature/chart/domain/command"
	res "gwi/assignment/feature/chart/domain/response"

	"github.com/google/uuid"
)

type ChartMapper struct{}

func (mapper *ChartMapper) ToDomainLayer(chart *ent.Chart) *res.ChartResponse {
	return &res.ChartResponse{
		Id:          chart.Id,
		AssetId:     chart.AssetId,
		XAxes:       chart.XAxes,
		YAxes:       chart.YAxes,
		Description: chart.Asset.Description,
		Points:      *mapper.pointListToDomainLayer(&chart.Points),
	}
}

func (mapper *ChartMapper) ToDomainLayerList(charts *[]ent.Chart) *[]res.ChartResponse {
	ret := []res.ChartResponse{}
	for _, chart := range *charts {
		ret = append(ret, *mapper.ToDomainLayer(&chart))
	}

	return &ret
}

func (mapper *ChartMapper) ToDomainLayerPaging(charts *[]ent.Chart, page int) *res.ChartPageResponse {
	return &res.ChartPageResponse{
		Page:    page,
		Results: *mapper.ToDomainLayerList(charts),
	}
}

func (mapper *ChartMapper) ToDataLayer(cmd *cmd.ChartCommand) *ent.Chart {
	assetEntity := &assetEnt.Asset{
		Description: cmd.Description,
	}
	assetEntity.Id = uuid.NewString()

	chartEntity := &ent.Chart{
		XAxes:   cmd.XAxes,
		YAxes:   cmd.YAxes,
		Asset:   *assetEntity,
		AssetId: assetEntity.Id,
	}
	chartEntity.Id = uuid.NewString()

	for _, chart := range cmd.Points {
		chartPointEntity := ent.ChartPoint{
			XValue:  chart.XValue,
			YValue:  chart.YValue,
			ChartId: chartEntity.Id,
		}
		chartPointEntity.Id = uuid.NewString()

		chartEntity.Points = append(chartEntity.Points, chartPointEntity)
	}

	return chartEntity
}

func (mapper *ChartMapper) pointToDomainLayer(point *ent.ChartPoint) *res.ChartPointResponse {
	return &res.ChartPointResponse{
		Id:     point.Id,
		XValue: point.XValue,
		YValue: point.YValue,
	}
}

func (mapper *ChartMapper) pointListToDomainLayer(points *[]ent.ChartPoint) *[]res.ChartPointResponse {
	ret := []res.ChartPointResponse{}
	for _, point := range *points {
		ret = append(ret, *mapper.pointToDomainLayer(&point))
	}

	return &ret
}
