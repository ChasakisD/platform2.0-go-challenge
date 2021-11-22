package usecase

import (
	domain "gwi/assignment/feature/audience/domain"
	cmd "gwi/assignment/feature/audience/domain/command"
	res "gwi/assignment/feature/audience/domain/response"
)

type AudienceUseCase struct {
	audienceMapper     domain.AudienceMapper
	audienceRepository domain.AudienceRepository
}

func Create(audienceMapper domain.AudienceMapper, audienceRepository domain.AudienceRepository) *AudienceUseCase {
	return &AudienceUseCase{
		audienceMapper:     audienceMapper,
		audienceRepository: audienceRepository,
	}
}

func (useCase *AudienceUseCase) GetAllAudiences(page int) (*res.AudiencePageResponse, error) {
	audiences, err := useCase.audienceRepository.GetAllAudiences(page)
	if err != nil {
		return nil, err
	}

	return useCase.audienceMapper.ToDomainLayerPaging(audiences, page), nil
}

func (useCase *AudienceUseCase) GetAudienceById(audienceId string) (*res.AudienceResponse, error) {
	audience, err := useCase.audienceRepository.GetAudienceById(audienceId)
	if err != nil {
		return nil, err
	}

	return useCase.audienceMapper.ToDomainLayer(audience), nil
}

func (useCase *AudienceUseCase) GetFavoriteAudiences(userId string, page int) (*res.AudiencePageResponse, error) {
	audiences, err := useCase.audienceRepository.GetFavoriteAudiences(userId, page)
	if err != nil {
		return nil, err
	}

	return useCase.audienceMapper.ToDomainLayerPaging(audiences, page), nil
}

func (useCase *AudienceUseCase) CreateAudience(command cmd.AudienceCommand) error {
	return useCase.audienceRepository.CreateAudience(useCase.audienceMapper.ToDataLayer(&command))
}

func (useCase *AudienceUseCase) UpdateAudience(audienceId string, command cmd.AudienceCommand) error {
	return useCase.audienceRepository.UpdateAudience(audienceId, useCase.audienceMapper.ToDataLayer(&command))
}

func (useCase *AudienceUseCase) DeleteAudience(audienceId string) error {
	return useCase.audienceRepository.DeleteAudience(audienceId)
}

func (useCase *AudienceUseCase) FavoriteAudience(userId, assetId string) error {
	return useCase.audienceRepository.FavoriteAudience(userId, assetId)
}

func (useCase *AudienceUseCase) UnfavoriteAudience(userId, assetId string) error {
	return useCase.audienceRepository.UnfavoriteAudience(userId, assetId)
}
