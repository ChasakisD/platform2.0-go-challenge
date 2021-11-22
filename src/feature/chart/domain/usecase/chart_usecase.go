package usecase

import (
	domain "gwi/assignment/feature/chart/domain"
	cmd "gwi/assignment/feature/chart/domain/command"
	res "gwi/assignment/feature/chart/domain/response"
)

type ChartUseCase struct {
	chartMapper     domain.ChartMapper
	chartRepository domain.ChartRepository
}

func Create(chartMapper domain.ChartMapper, chartRepository domain.ChartRepository) *ChartUseCase {
	return &ChartUseCase{
		chartMapper:     chartMapper,
		chartRepository: chartRepository,
	}
}

func (useCase *ChartUseCase) GetAllCharts(page int) (*res.ChartPageResponse, error) {
	charts, err := useCase.chartRepository.GetAllCharts(page)
	if err != nil {
		return nil, err
	}

	return useCase.chartMapper.ToDomainLayerPaging(charts, page), nil
}

func (useCase *ChartUseCase) GetChartById(chartId string) (*res.ChartResponse, error) {
	chart, err := useCase.chartRepository.GetChartById(chartId)
	if err != nil {
		return nil, err
	}

	return useCase.chartMapper.ToDomainLayer(chart), nil
}

func (useCase *ChartUseCase) GetFavoriteCharts(userId string, page int) (*res.ChartPageResponse, error) {
	charts, err := useCase.chartRepository.GetFavoriteCharts(userId, page)
	if err != nil {
		return nil, err
	}

	return useCase.chartMapper.ToDomainLayerPaging(charts, page), nil
}

func (useCase *ChartUseCase) CreateChart(command cmd.ChartCommand) error {
	return useCase.chartRepository.CreateChart(useCase.chartMapper.ToDataLayer(&command))
}

func (useCase *ChartUseCase) UpdateChart(chartId string, command cmd.ChartCommand) error {
	return useCase.chartRepository.UpdateChart(chartId, useCase.chartMapper.ToDataLayer(&command))
}

func (useCase *ChartUseCase) DeleteChart(chartId string) error {
	return useCase.chartRepository.DeleteChart(chartId)
}

func (useCase *ChartUseCase) FavoriteChart(userId, assetId string) error {
	return useCase.chartRepository.FavoriteChart(userId, assetId)
}

func (useCase *ChartUseCase) UnfavoriteChart(userId, assetId string) error {
	return useCase.chartRepository.UnfavoriteChart(userId, assetId)
}
