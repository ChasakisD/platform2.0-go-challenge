package insight

import (
	"gwi/assignment/core/data/http"
	"gwi/assignment/feature/insight/api"
	"gwi/assignment/feature/insight/data/mapper"
	"gwi/assignment/feature/insight/data/repository"
	"gwi/assignment/feature/insight/domain/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func Initialize(router *chi.Mux) {
	insightMapper := &mapper.InsightMapper{}
	insightRepository := &repository.InsightRepository{}
	insightUseCase := usecase.Create(insightMapper, insightRepository)
	insightController := api.InsightController{InsightUseCase: insightUseCase}

	//Public Endpoints
	router.Group(func(router chi.Router) {
		router.Route("/insight", func(router chi.Router) {
			router.Get("/", insightController.GetAllInsights)
			router.Get("/{insightId}", insightController.GetInsightById)
		})
	})

	//Protected Endpoints
	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(http.TokenAuth))
		router.Use(jwtauth.Authenticator)

		router.Post("/insight", insightController.CreateInsight)

		router.Get("/insight/favorite", insightController.GetFavoriteInsights)
		router.Put("/insight/{insightId}", insightController.UpdateInsight)
		router.Delete("/insight/{insightId}", insightController.DeleteInsight)

		router.Route("/insight/{insightId}/favorite", func(router chi.Router) {
			router.Post("/", insightController.FavoriteInsight)
			router.Delete("/", insightController.UnfavoriteInsight)
		})
	})
}
