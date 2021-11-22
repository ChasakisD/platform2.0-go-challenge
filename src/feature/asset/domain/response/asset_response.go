package response

import (
	audienceRes "gwi/assignment/feature/audience/domain/response"
	chartRes "gwi/assignment/feature/chart/domain/response"
	insightRes "gwi/assignment/feature/insight/domain/response"
)

type AssetResponse struct {
	Audiences audienceRes.AudiencePageResponse `json:"audiences"`
	Charts    chartRes.ChartPageResponse       `json:"charts"`
	Insights  insightRes.InsightPageResponse   `json:"insights"`
}
