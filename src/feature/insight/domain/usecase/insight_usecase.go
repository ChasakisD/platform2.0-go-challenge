package usecase

import (
	domain "gwi/assignment/feature/insight/domain"
	cmd "gwi/assignment/feature/insight/domain/command"
	res "gwi/assignment/feature/insight/domain/response"
)

type InsightUseCase struct {
	insightMapper     domain.InsightMapper
	insightRepository domain.InsightRepository
}

func Create(insightMapper domain.InsightMapper, insightRepository domain.InsightRepository) *InsightUseCase {
	return &InsightUseCase{
		insightMapper:     insightMapper,
		insightRepository: insightRepository,
	}
}

func (useCase *InsightUseCase) GetAllInsights(page int) (*res.InsightPageResponse, error) {
	audiences, err := useCase.insightRepository.GetAllInsights(page)
	if err != nil {
		return nil, err
	}

	return useCase.insightMapper.ToDomainLayerPaging(audiences, page), nil
}

func (useCase *InsightUseCase) GetInsightById(audienceId string) (*res.InsightResponse, error) {
	audience, err := useCase.insightRepository.GetInsightById(audienceId)
	if err != nil {
		return nil, err
	}

	return useCase.insightMapper.ToDomainLayer(audience), nil
}

func (useCase *InsightUseCase) GetFavoriteInsights(userId string, page int) (*res.InsightPageResponse, error) {
	audiences, err := useCase.insightRepository.GetFavoriteInsights(userId, page)
	if err != nil {
		return nil, err
	}

	return useCase.insightMapper.ToDomainLayerPaging(audiences, page), nil
}

func (useCase *InsightUseCase) CreateInsight(command cmd.InsightCommand) error {
	return useCase.insightRepository.CreateInsight(useCase.insightMapper.ToDataLayer(&command))
}

func (useCase *InsightUseCase) UpdateInsight(audienceId string, command cmd.InsightCommand) error {
	return useCase.insightRepository.UpdateInsight(audienceId, useCase.insightMapper.ToDataLayer(&command))
}

func (useCase *InsightUseCase) DeleteInsight(audienceId string) error {
	return useCase.insightRepository.DeleteInsight(audienceId)
}

func (useCase *InsightUseCase) FavoriteInsight(userId, assetId string) error {
	return useCase.insightRepository.FavoriteInsight(userId, assetId)
}

func (useCase *InsightUseCase) UnfavoriteInsight(userId, assetId string) error {
	return useCase.insightRepository.UnfavoriteInsight(userId, assetId)
}
