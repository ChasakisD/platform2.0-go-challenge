package response

type AudiencePageResponse struct {
	Page    int                `json:"page"`
	Results []AudienceResponse `json:"results"`
}

type AudienceResponse struct {
	Id                   string                   `json:"id"`
	AssetId              string                   `json:"assetId"`
	Description          string                   `json:"description"`
	DescriptionFormatted string                   `json:"descriptionFormatted"`
	Gender               string                   `json:"gender"`
	BirthCountry         string                   `json:"birthCountry"`
	AgeGroupMin          int                      `json:"ageGroupMin"`
	AgeGroupMax          int                      `json:"ageGroupMax"`
	StatType             AudienceStatTypeResponse `json:"statType"`
	StatTypeValue        float64                  `json:"statTypeValue"`
}

type AudienceStatTypeResponse struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	TitleFormatted string `json:"titleFormatted"`
}
