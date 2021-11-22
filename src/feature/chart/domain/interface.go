package domain

import (
	ent "gwi/assignment/feature/chart/data/entity"
	cmd "gwi/assignment/feature/chart/domain/command"
	res "gwi/assignment/feature/chart/domain/response"
)

type ChartRepository interface {
	GetAllCharts(page int) (*[]ent.Chart, error)
	GetChartById(chartId string) (*ent.Chart, error)
	GetFavoriteCharts(userId string, page int) (*[]ent.Chart, error)

	CreateChart(chart *ent.Chart) error
	UpdateChart(chartId string, chart *ent.Chart) error
	DeleteChart(chartId string) error

	FavoriteChart(userId string, chartId string) error
	UnfavoriteChart(userId string, chartId string) error
}

type ChartMapper interface {
	ToDataLayer(chart *cmd.ChartCommand) *ent.Chart
	ToDomainLayer(chart *ent.Chart) (ret *res.ChartResponse)
	ToDomainLayerList(charts *[]ent.Chart) (ret *[]res.ChartResponse)
	ToDomainLayerPaging(charts *[]ent.Chart, page int) *res.ChartPageResponse
}
