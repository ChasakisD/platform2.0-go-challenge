package response

type InsightPageResponse struct {
	Page    int               `json:"page"`
	Results []InsightResponse `json:"results"`
}

type InsightResponse struct {
	Id          string `json:"id"`
	AssetId     string `json:"assetId"`
	Description string `json:"description"`
}
