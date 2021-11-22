package entity

type BaseEntity struct {
	Id string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
}
