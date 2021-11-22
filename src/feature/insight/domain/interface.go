package domain

import (
	ent "gwi/assignment/feature/insight/data/entity"
	cmd "gwi/assignment/feature/insight/domain/command"
	res "gwi/assignment/feature/insight/domain/response"
)

type InsightRepository interface {
	GetAllInsights(page int) (*[]ent.Insight, error)
	GetInsightById(insightId string) (*ent.Insight, error)
	GetFavoriteInsights(userId string, page int) (*[]ent.Insight, error)

	CreateInsight(insight *ent.Insight) error
	UpdateInsight(insightId string, audience *ent.Insight) error
	DeleteInsight(insightId string) error

	FavoriteInsight(userId string, insightId string) error
	UnfavoriteInsight(userId string, insightId string) error
}

type InsightMapper interface {
	ToDataLayer(insight *cmd.InsightCommand) *ent.Insight
	ToDomainLayer(insight *ent.Insight) (ret *res.InsightResponse)
	ToDomainLayerList(insights *[]ent.Insight) (ret *[]res.InsightResponse)
	ToDomainLayerPaging(insights *[]ent.Insight, page int) *res.InsightPageResponse
}
