package domain

import (
	ent "gwi/assignment/feature/audience/data/entity"
	cmd "gwi/assignment/feature/audience/domain/command"
	res "gwi/assignment/feature/audience/domain/response"
)

type AudienceRepository interface {
	GetAllAudiences(page int) (*[]ent.Audience, error)
	GetAudienceById(audienceId string) (*ent.Audience, error)
	GetFavoriteAudiences(userId string, page int) (*[]ent.Audience, error)

	CreateAudience(audience *ent.Audience) error
	UpdateAudience(audienceId string, audience *ent.Audience) error
	DeleteAudience(audienceId string) error

	FavoriteAudience(userId string, audienceId string) error
	UnfavoriteAudience(userId string, audienceId string) error
}

type AudienceMapper interface {
	ToDataLayer(audience *cmd.AudienceCommand) *ent.Audience
	ToDomainLayer(audience *ent.Audience) (ret *res.AudienceResponse)
	ToDomainLayerList(audiences *[]ent.Audience) (ret *[]res.AudienceResponse)
	ToDomainLayerPaging(audiences *[]ent.Audience, page int) *res.AudiencePageResponse
}
