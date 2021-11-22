package audience

import (
	"gwi/assignment/core/data/http"
	"gwi/assignment/feature/audience/api"
	"gwi/assignment/feature/audience/data/mapper"
	"gwi/assignment/feature/audience/data/repository"
	"gwi/assignment/feature/audience/domain/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func Initialize(router *chi.Mux) {
	audienceMapper := &mapper.AudienceMapper{}
	audienceRepository := &repository.AudienceRepository{}
	audienceUseCase := usecase.Create(audienceMapper, audienceRepository)
	audienceController := api.AudienceController{AudienceUseCase: audienceUseCase}

	//Public Endpoints
	router.Group(func(router chi.Router) {
		router.Route("/audience", func(router chi.Router) {
			router.Get("/", audienceController.GetAllAudiences)
			router.Get("/{audienceId}", audienceController.GetAudienceById)
		})
	})

	//Protected Endpoints
	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(http.TokenAuth))
		router.Use(jwtauth.Authenticator)

		router.Post("/audience", audienceController.CreateAudience)

		router.Get("/audience/favorite", audienceController.GetFavoriteAudiences)
		router.Put("/audience/{audienceId}", audienceController.UpdateAudience)
		router.Delete("/audience/{audienceId}", audienceController.DeleteAudience)

		router.Route("/audience/{audienceId}/favorite", func(router chi.Router) {
			router.Post("/", audienceController.FavoriteAudience)
			router.Delete("/", audienceController.UnfavoriteAudience)
		})
	})
}
