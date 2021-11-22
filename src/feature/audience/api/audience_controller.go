package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	apiCore "gwi/assignment/core/api"
	coreRes "gwi/assignment/core/domain/response"
	cmd "gwi/assignment/feature/audience/domain/command"
	res "gwi/assignment/feature/audience/domain/response"
	useCase "gwi/assignment/feature/audience/domain/usecase"

	"github.com/go-chi/chi/v5"
)

type AudienceController struct {
	AudienceUseCase *useCase.AudienceUseCase
}

func dummySwaggerModels() {
	print(coreRes.ErrorResponse{}.Error)
	print(res.AudiencePageResponse{}.Page)
}

// @Summary Get audiences
// @Description Returns all the application audiences
// @Tags Audience
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Success 200 {object} res.AudiencePageResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 500 {object} coreRes.ErrorResponse
// @Router /audience [get]
func (uc *AudienceController) GetAllAudiences(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		apiCore.RenderError(w, r, apiCore.ErrParamNotValid)
		return
	}

	audiences, err := uc.AudienceUseCase.GetAllAudiences(page)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, audiences)
}

// @Summary Get audience
// @Description Returns audience by id
// @Tags Audience
// @Accept json
// @Produce json
// @Param audienceId path string false "audienceId"
// @Success 200 {object} res.AudienceResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 500 {object} coreRes.ErrorResponse
// @Router /audience/{audienceId} [get]
func (uc *AudienceController) GetAudienceById(w http.ResponseWriter, r *http.Request) {
	audienceId, err := getAudienceId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	audience, err := uc.AudienceUseCase.GetAudienceById(audienceId)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, audience)
}

// @Summary Get favorite audiences
// @Description Get user's favorite audiences
// @Tags Audience
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Success 200 {object} res.AudiencePageResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /audience/favorite [get]
func (uc *AudienceController) GetFavoriteAudiences(w http.ResponseWriter, r *http.Request) {
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

	audiences, err := uc.AudienceUseCase.GetFavoriteAudiences(userId, page)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, audiences)
}

// @Summary Create audience
// @Description Creates an audience
// @Tags Audience
// @Accept json
// @Produce json
// @Param command body cmd.AudienceCommand false "audienceId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /audience [post]
func (uc *AudienceController) CreateAudience(w http.ResponseWriter, r *http.Request) {
	command := &cmd.AudienceCommand{}
	if err := json.NewDecoder(r.Body).Decode(command); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := command.Validate(); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.AudienceUseCase.CreateAudience(*command); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

// @Summary Update audience
// @Description Updates an audience
// @Tags Audience
// @Accept json
// @Produce json
// @Param audienceId path string false "audienceId"
// @Param command body cmd.AudienceCommand false "audienceId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /audience/{audienceId} [put]
func (uc *AudienceController) UpdateAudience(w http.ResponseWriter, r *http.Request) {
	audienceId, err := getAudienceId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	command := &cmd.AudienceCommand{}
	if err = json.NewDecoder(r.Body).Decode(command); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err = command.Validate(); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.AudienceUseCase.UpdateAudience(audienceId, *command); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

// @Summary Delete audience
// @Description Deletes an audience
// @Tags Audience
// @Accept json
// @Produce json
// @Param audienceId path string false "audienceId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /audience/{audienceId} [delete]
func (uc *AudienceController) DeleteAudience(w http.ResponseWriter, r *http.Request) {
	audienceId, err := getAudienceId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.AudienceUseCase.DeleteAudience(audienceId); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

// @Summary Favorite audience
// @Description Favorites an audience
// @Tags Audience
// @Accept json
// @Produce json
// @Param audienceId path string false "audienceId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /audience/{audienceId}/favorite [post]
func (uc *AudienceController) FavoriteAudience(w http.ResponseWriter, r *http.Request) {
	userId, err := apiCore.GetUserIdFromContext(r.Context())
	if err != nil {
		apiCore.RenderUnauthorized(r)
		return
	}

	audienceId, err := getAudienceId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.AudienceUseCase.FavoriteAudience(userId, audienceId); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

// @Summary Unfavorite audience
// @Description Unfavorites an audience
// @Tags Audience
// @Accept json
// @Produce json
// @Param audienceId path string false "audienceId"
// @Success 200
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 401
// @Failure 500 {object} coreRes.ErrorResponse
// @Security ApiKeyAuth
// @Router /audience/{audienceId}/favorite [delete]
func (uc *AudienceController) UnfavoriteAudience(w http.ResponseWriter, r *http.Request) {
	userId, err := apiCore.GetUserIdFromContext(r.Context())
	if err != nil {
		apiCore.RenderUnauthorized(r)
		return
	}

	audienceId, err := getAudienceId(r)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	if err := uc.AudienceUseCase.UnfavoriteAudience(userId, audienceId); err != nil {
		apiCore.RenderError(w, r, err)
		return
	}
}

func getAudienceId(r *http.Request) (string, error) {
	audienceId := chi.URLParam(r, "audienceId")
	if len(audienceId) <= 0 {
		return "", apiCore.ErrParamNotValid
	}

	return audienceId, nil
}
