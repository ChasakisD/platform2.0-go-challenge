package response

type ChartPageResponse struct {
	Page    int             `json:"page"`
	Results []ChartResponse `json:"results"`
}

type ChartResponse struct {
	Id          string               `json:"id"`
	AssetId     string               `json:"assetId"`
	Description string               `json:"description"`
	XAxes       string               `json:"xAxes"`
	YAxes       string               `json:"yAxes"`
	Points      []ChartPointResponse `json:"points"`
}

type ChartPointResponse struct {
	Id     string  `json:"id"`
	XValue float64 `json:"xValue"`
	YValue float64 `json:"yValue"`
}
