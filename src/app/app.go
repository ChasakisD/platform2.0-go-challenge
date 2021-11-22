package app

import (
	"net/http"

	"gwi/assignment/core"
	_ "gwi/assignment/docs"
	"gwi/assignment/environment"
	"gwi/assignment/feature/asset"
	"gwi/assignment/feature/audience"
	"gwi/assignment/feature/chart"
	"gwi/assignment/feature/insight"
	"gwi/assignment/feature/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Run() {
	config := environment.LoadConfig("config.yml")

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	initSwagger(router)
	initModules(router, config)

	startServer(router, config)
}

func initSwagger(router *chi.Mux) {
	router.Mount("/swagger", httpSwagger.Handler())
}

func initModules(router *chi.Mux, config environment.Config) {
	core.Initialize(config.Database.ConnectionString, config.JWT.Secret)

	asset.Initialize(router)
	audience.Initialize(router)
	chart.Initialize(router)
	insight.Initialize(router)
	user.Initialize(router)
}

func startServer(router *chi.Mux, config environment.Config) {
	http.ListenAndServe(config.Server.Address, router)
}
