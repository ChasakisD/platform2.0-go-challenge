package usecase

import (
	"errors"
	"sync"

	audienceUC "gwi/assignment/feature/audience/domain/usecase"
	chartUC "gwi/assignment/feature/chart/domain/usecase"
	insightUC "gwi/assignment/feature/insight/domain/usecase"

	res "gwi/assignment/feature/asset/domain/response"
	audienceRes "gwi/assignment/feature/audience/domain/response"
	chartRes "gwi/assignment/feature/chart/domain/response"
	insightRes "gwi/assignment/feature/insight/domain/response"
)

type AssetUseCase struct {
	audienceUseCase audienceUC.AudienceUseCase
	chartUseCase    chartUC.ChartUseCase
	insightUseCase  insightUC.InsightUseCase
}

func Create(audienceUseCase audienceUC.AudienceUseCase, chartUseCase chartUC.ChartUseCase, insightUseCase insightUC.InsightUseCase) *AssetUseCase {
	return &AssetUseCase{
		audienceUseCase: audienceUseCase,
		chartUseCase:    chartUseCase,
		insightUseCase:  insightUseCase,
	}
}

func (useCase *AssetUseCase) GetAllAssets(page int) (*res.AssetResponse, error) {
	assetResponse := &res.AssetResponse{}
	var errorStrings [3]string

	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	go useCase.getAllAudiences(page, &assetResponse.Audiences, &errorStrings[0], &waitGroup)
	go useCase.getAllCharts(page, &assetResponse.Charts, &errorStrings[1], &waitGroup)
	go useCase.getAllInsights(page, &assetResponse.Insights, &errorStrings[2], &waitGroup)

	waitGroup.Wait()

	for _, err := range errorStrings {
		if len(err) > 0 {
			return nil, errors.New(err)
		}
	}

	return assetResponse, nil
}

func (useCase *AssetUseCase) GetFavoriteAssets(userId string, page int) (*res.AssetResponse, error) {
	assetResponse := &res.AssetResponse{}
	var errorStrings [3]string

	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	go useCase.getFavoriteAudiences(userId, page, &assetResponse.Audiences, &errorStrings[0], &waitGroup)
	go useCase.getFavoriteCharts(userId, page, &assetResponse.Charts, &errorStrings[1], &waitGroup)
	go useCase.getFavoriteInsights(userId, page, &assetResponse.Insights, &errorStrings[2], &waitGroup)

	waitGroup.Wait()

	for _, err := range errorStrings {
		if len(err) > 0 {
			return nil, errors.New(err)
		}
	}

	return assetResponse, nil
}

func (useCase *AssetUseCase) getAllAudiences(page int, audiences *audienceRes.AudiencePageResponse, errorString *string, waitGroup *sync.WaitGroup) {
	audiencePage, err := useCase.audienceUseCase.GetAllAudiences(page)
	if audiencePage != nil {
		*audiences = *audiencePage
	}
	if err != nil {
		*errorString = err.Error()
	}
	waitGroup.Done()
}

func (useCase *AssetUseCase) getAllCharts(page int, charts *chartRes.ChartPageResponse, errorString *string, waitGroup *sync.WaitGroup) {
	chartPage, err := useCase.chartUseCase.GetAllCharts(page)
	if chartPage != nil {
		*charts = *chartPage
	}
	if err != nil {
		*errorString = err.Error()
	}
	waitGroup.Done()
}

func (useCase *AssetUseCase) getAllInsights(page int, insights *insightRes.InsightPageResponse, errorString *string, waitGroup *sync.WaitGroup) {
	insightPage, err := useCase.insightUseCase.GetAllInsights(page)
	if insightPage != nil {
		*insights = *insightPage
	}
	if err != nil {
		*errorString = err.Error()
	}
	waitGroup.Done()
}

func (useCase *AssetUseCase) getFavoriteAudiences(userId string, page int, audiences *audienceRes.AudiencePageResponse, errorString *string, waitGroup *sync.WaitGroup) {
	audiencePage, err := useCase.audienceUseCase.GetFavoriteAudiences(userId, page)
	if audiencePage != nil {
		*audiences = *audiencePage
	}
	if err != nil {
		*errorString = err.Error()
	}
	waitGroup.Done()
}

func (useCase *AssetUseCase) getFavoriteCharts(userId string, page int, charts *chartRes.ChartPageResponse, errorString *string, waitGroup *sync.WaitGroup) {
	chartPage, err := useCase.chartUseCase.GetFavoriteCharts(userId, page)
	if chartPage != nil {
		*charts = *chartPage
	}
	if err != nil {
		*errorString = err.Error()
	}
	waitGroup.Done()
}

func (useCase *AssetUseCase) getFavoriteInsights(userId string, page int, insights *insightRes.InsightPageResponse, errorString *string, waitGroup *sync.WaitGroup) {
	insightPage, err := useCase.insightUseCase.GetFavoriteInsights(userId, page)
	if insightPage != nil {
		*insights = *insightPage
	}
	if err != nil {
		*errorString = err.Error()
	}
	waitGroup.Done()
}
