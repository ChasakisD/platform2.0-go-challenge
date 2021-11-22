package api

import (
	"net/http"
	"strconv"

	apiCore "gwi/assignment/core/api"
	coreRes "gwi/assignment/core/domain/response"
	res "gwi/assignment/feature/asset/domain/response"
	useCase "gwi/assignment/feature/asset/domain/usecase"
)

type AssetController struct {
	AssetUseCase *useCase.AssetUseCase
}

func dummySwaggerModels() {
	print(coreRes.ErrorResponse{}.Error)
	print(res.AssetResponse{}.Audiences.Page)
}

// @Summary Get assets
// @Description Returns all the application assets
// @Tags Asset
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Success 200 {object} res.AssetResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 500 {object} coreRes.ErrorResponse
// @Router /asset [get]
func (uc *AssetController) GetAllAssets(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		apiCore.RenderError(w, r, apiCore.ErrParamNotValid)
		return
	}

	assets, err := uc.AssetUseCase.GetAllAssets(page)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, assets)
}

// @Summary Get favorite assets
// @Description Get user's favorite assets
// @Tags Asset
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Success 200 {object} res.AssetResponse
// @Failure 400 {object} coreRes.ErrorResponse
// @Failure 500 {object} coreRes.ErrorResponse
// @Failure 401
// @Security ApiKeyAuth
// @Router /asset/favorite [get]
func (uc *AssetController) GetFavoriteAssets(w http.ResponseWriter, r *http.Request) {
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

	assets, err := uc.AssetUseCase.GetFavoriteAssets(userId, page)
	if err != nil {
		apiCore.RenderError(w, r, err)
		return
	}

	apiCore.RenderOk(w, r, assets)
}
