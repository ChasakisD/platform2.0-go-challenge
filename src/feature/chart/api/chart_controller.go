package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	apiCore "gwi/assignment/core/api"
	coreRes "gwi/assignment/core/domain/response"
	cmd "gwi/assignment/feature/chart/domain/command"
	res "gwi/assignment/feature/chart/domain/response"
	useCase "gwi/assignment/feature/chart/domain/usecase"

	"github.com/go-chi/chi/v5"
)

type ChartController struct {
	ChartUseCase *useCase.ChartUseCase
}

func dummySwaggerModels() {
	print(coreRes.ErrorResponse{}.Error)
	print(res.ChartPageResponse{}.Page)
}

// @Summary Get charts
// @Description Returns all the application charts
// @Tags Chart
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Success 200 {object} res.ChartPageResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 500 {object} coreRes.ErrorResponse
// @Router /chart [get]
func (uc *ChartController) GetAllCharts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		apiCore.RenderError(w, r, apiCore.ErrParamNotValid)
		return
	}

	charts, err := uc.ChartUseCase.GetAllCharts(page)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, charts)
}

// @Summary Get chart
// @Description Returns chart by id
// @Tags Chart
// @Accept json
// @Produce json
// @Param chartId path string false "chartId"
// @Success 200 {object} res.ChartResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 500 {object} coreRes.ErrorResponse
// @Router /chart/{chartId} [get]
func (uc *ChartController) GetChartById(w http.ResponseWriter, r *http.Request) {
	chartId, err := getChartId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	chart, err := uc.ChartUseCase.GetChartById(chartId)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, chart)
}

// @Summary Get favorite charts
// @Description Get user's favorite charts
// @Tags Chart
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Success 200 {object} res.ChartPageResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /chart/favorite [get]
func (uc *ChartController) GetFavoriteCharts(w http.ResponseWriter, r *http.Request) {
	userId, err := apiCore.GetUserIdFromContext(r.Context())
	if err != nil {
		apiCore.RenderUnauthorized(r)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		apiCore.RenderError(w, r, apiCore.ErrParamNotValid)
		return
	}

	charts, err := uc.ChartUseCase.GetFavoriteCharts(userId, page)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, charts)
}

// @Summary Create chart
// @Description Creates an chart
// @Tags Chart
// @Accept json
// @Produce json
// @Param command body cmd.ChartCommand false "chartId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /chart [post]
func (uc *ChartController) CreateChart(w http.ResponseWriter, r *http.Request) {
	command := &cmd.ChartCommand{}
	if err := json.NewDecoder(r.Body).Decode(command); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := command.Validate(); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.ChartUseCase.CreateChart(*command); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

// @Summary Update chart
// @Description Updates an chart
// @Tags Chart
// @Accept json
// @Produce json
// @Param chartId path string false "chartId"
// @Param command body cmd.ChartCommand false "chartId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /chart/{chartId} [put]
func (uc *ChartController) UpdateChart(w http.ResponseWriter, r *http.Request) {
	chartId, err := getChartId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	command := &cmd.ChartCommand{}
	if err = json.NewDecoder(r.Body).Decode(command); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err = command.Validate(); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.ChartUseCase.UpdateChart(chartId, *command); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

// @Summary Delete chart
// @Description Deletes an chart
// @Tags Chart
// @Accept json
// @Produce json
// @Param chartId path string false "chartId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /chart/{chartId} [delete]
func (uc *ChartController) DeleteChart(w http.ResponseWriter, r *http.Request) {
	chartId, err := getChartId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.ChartUseCase.DeleteChart(chartId); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

// @Summary Favorite chart
// @Description Favorites an chart
// @Tags Chart
// @Accept json
// @Produce json
// @Param chartId path string false "chartId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /chart/{chartId}/favorite [post]
func (uc *ChartController) FavoriteChart(w http.ResponseWriter, r *http.Request) {
	userId, err := apiCore.GetUserIdFromContext(r.Context())
	if err != nil {
		apiCore.RenderUnauthorized(r)
		return
	}

	chartId, err := getChartId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.ChartUseCase.FavoriteChart(userId, chartId); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

// @Summary Unfavorite chart
// @Description Unfavorites an chart
// @Tags Chart
// @Accept json
// @Produce json
// @Param chartId path string false "chartId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /chart/{chartId}/favorite [delete]
func (uc *ChartController) UnfavoriteChart(w http.ResponseWriter, r *http.Request) {
	userId, err := apiCore.GetUserIdFromContext(r.Context())
	if err != nil {
		apiCore.RenderUnauthorized(r)
		return
	}

	chartId, err := getChartId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.ChartUseCase.UnfavoriteChart(userId, chartId); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

func getChartId(r *http.Request) (string, error) {
	chartId := chi.URLParam(r, "chartId")
	if len(chartId) <= 0 {
		return "", apiCore.ErrParamNotValid
	}

	return chartId, nil
}
