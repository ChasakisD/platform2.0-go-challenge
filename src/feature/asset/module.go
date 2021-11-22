package asset

import (
	"gwi/assignment/core/data/http"
	"gwi/assignment/feature/asset/api"
	"gwi/assignment/feature/asset/domain/usecase"

	audienceMapper "gwi/assignment/feature/audience/data/mapper"
	audienceRepo "gwi/assignment/feature/audience/data/repository"
	audienceUc "gwi/assignment/feature/audience/domain/usecase"

	chartMapper "gwi/assignment/feature/chart/data/mapper"
	chartRepo "gwi/assignment/feature/chart/data/repository"
	chartUc "gwi/assignment/feature/chart/domain/usecase"

	insightMapper "gwi/assignment/feature/insight/data/mapper"
	insightRepo "gwi/assignment/feature/insight/data/repository"
	insightUc "gwi/assignment/feature/insight/domain/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func Initialize(router *chi.Mux) {
	audienceMapper := &audienceMapper.AudienceMapper{}
	audienceRepository := &audienceRepo.AudienceRepository{}
	audienceUseCase := audienceUc.Create(audienceMapper, audienceRepository)

	chartMapper := &chartMapper.ChartMapper{}
	chartRepository := &chartRepo.ChartRepository{}
	chartUseCase := chartUc.Create(chartMapper, chartRepository)

	insightMapper := &insightMapper.InsightMapper{}
	insightRepository := &insightRepo.InsightRepository{}
	insightUseCase := insightUc.Create(insightMapper, insightRepository)

	assetUseCase := usecase.Create(*audienceUseCase, *chartUseCase, *insightUseCase)
	assetController := api.AssetController{AssetUseCase: assetUseCase}

	//Public Endpoints
	router.Group(func(router chi.Router) {
		router.Route("/asset", func(router chi.Router) {
			router.Get("/", assetController.GetAllAssets)
		})
	})

	//Protected Endpoints
	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(http.TokenAuth))
		router.Use(jwtauth.Authenticator)

		router.Get("/asset/favorite", assetController.GetFavoriteAssets)
	})
}
