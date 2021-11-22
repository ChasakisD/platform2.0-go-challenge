package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	apiCore "gwi/assignment/core/api"
	coreRes "gwi/assignment/core/domain/response"
	cmd "gwi/assignment/feature/insight/domain/command"
	res "gwi/assignment/feature/insight/domain/response"
	useCase "gwi/assignment/feature/insight/domain/usecase"

	"github.com/go-chi/chi/v5"
)

type InsightController struct {
	InsightUseCase *useCase.InsightUseCase
}

func dummySwaggerModels() {
	print(coreRes.ErrorResponse{}.Error)
	print(res.InsightPageResponse{}.Page)
}

// @Summary Get insights
// @Description Returns all the application insights
// @Tags Insight
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Success 200 {object} res.InsightPageResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 500 {object} coreRes.ErrorResponse
// @Router /insight [get]
func (uc *InsightController) GetAllInsights(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		apiCore.RenderError(w, r, apiCore.ErrParamNotValid)
		return
	}

	insights, err := uc.InsightUseCase.GetAllInsights(page)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, insights)
}

// @Summary Get insight
// @Description Returns insight by id
// @Tags Insight
// @Accept json
// @Produce json
// @Param insightId path string false "insightId"
// @Success 200 {object} res.InsightResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 500 {object} coreRes.ErrorResponse
// @Router /insight/{insightId} [get]
func (uc *InsightController) GetInsightById(w http.ResponseWriter, r *http.Request) {
	insightId, err := getInsightId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	insight, err := uc.InsightUseCase.GetInsightById(insightId)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, insight)
}

// @Summary Get favorite insights
// @Description Get user's favorite insights
// @Tags Insight
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Success 200 {object} res.InsightPageResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /insight/favorite [get]
func (uc *InsightController) GetFavoriteInsights(w http.ResponseWriter, r *http.Request) {
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

	insights, err := uc.InsightUseCase.GetFavoriteInsights(userId, page)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, insights)
}

// @Summary Create insight
// @Description Creates an insight
// @Tags Insight
// @Accept json
// @Produce json
// @Param command body cmd.InsightCommand false "insightId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /insight [post]
func (uc *InsightController) CreateInsight(w http.ResponseWriter, r *http.Request) {
	command := &cmd.InsightCommand{}
	if err := json.NewDecoder(r.Body).Decode(command); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := command.Validate(); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.InsightUseCase.CreateInsight(*command); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

// @Summary Update insight
// @Description Updates an insight
// @Tags Insight
// @Accept json
// @Produce json
// @Param insightId path string false "insightId"
// @Param command body cmd.InsightCommand false "insightId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /insight/{insightId} [put]
func (uc *InsightController) UpdateInsight(w http.ResponseWriter, r *http.Request) {
	insightId, err := getInsightId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	command := &cmd.InsightCommand{}
	if err = json.NewDecoder(r.Body).Decode(command); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err = command.Validate(); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.InsightUseCase.UpdateInsight(insightId, *command); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

// @Summary Delete insight
// @Description Deletes an insight
// @Tags Insight
// @Accept json
// @Produce json
// @Param insightId path string false "insightId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /insight/{insightId} [delete]
func (uc *InsightController) DeleteInsight(w http.ResponseWriter, r *http.Request) {
	insightId, err := getInsightId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.InsightUseCase.DeleteInsight(insightId); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

// @Summary Favorite insight
// @Description Favorites an insight
// @Tags Insight
// @Accept json
// @Produce json
// @Param insightId path string false "insightId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /insight/{insightId}/favorite [post]
func (uc *InsightController) FavoriteInsight(w http.ResponseWriter, r *http.Request) {
	userId, err := apiCore.GetUserIdFromContext(r.Context())
	if err != nil {
		apiCore.RenderUnauthorized(r)
		return
	}

	insightId, err := getInsightId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.InsightUseCase.FavoriteInsight(userId, insightId); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

// @Summary Unfavorite insight
// @Description Unfavorites an insight
// @Tags Insight
// @Accept json
// @Produce json
// @Param insightId path string false "insightId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /insight/{insightId}/favorite [delete]
func (uc *InsightController) UnfavoriteInsight(w http.ResponseWriter, r *http.Request) {
	userId, err := apiCore.GetUserIdFromContext(r.Context())
	if err != nil {
		apiCore.RenderUnauthorized(r)
		return
	}

	insightId, err := getInsightId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.InsightUseCase.UnfavoriteInsight(userId, insightId); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

func getInsightId(r *http.Request) (string, error) {
	insightId := chi.URLParam(r, "insightId")
	if len(insightId) <= 0 {
		return "", apiCore.ErrParamNotValid
	}

	return insightId, nil
}
