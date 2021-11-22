package chart

import (
	"gwi/assignment/core/data/http"
	"gwi/assignment/feature/chart/api"
	"gwi/assignment/feature/chart/data/mapper"
	"gwi/assignment/feature/chart/data/repository"
	"gwi/assignment/feature/chart/domain/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func Initialize(router *chi.Mux) {
	chartMapper := &mapper.ChartMapper{}
	chartRepository := &repository.ChartRepository{}
	chartUseCase := usecase.Create(chartMapper, chartRepository)
	chartController := api.ChartController{ChartUseCase: chartUseCase}

	//Public Endpoints
	router.Group(func(router chi.Router) {
		router.Route("/chart", func(router chi.Router) {
			router.Get("/", chartController.GetAllCharts)
			router.Get("/{chartId}", chartController.GetChartById)
		})
	})

	//Protected Endpoints
	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(http.TokenAuth))
		router.Use(jwtauth.Authenticator)

		router.Post("/chart", chartController.CreateChart)

		router.Get("/chart/favorite", chartController.GetFavoriteCharts)
		router.Put("/chart/{chartId}", chartController.UpdateChart)
		router.Delete("/chart/{chartId}", chartController.DeleteChart)

		router.Route("/chart/{chartId}/favorite", func(router chi.Router) {
			router.Post("/", chartController.FavoriteChart)
			router.Delete("/", chartController.UnfavoriteChart)
		})
	})
}
