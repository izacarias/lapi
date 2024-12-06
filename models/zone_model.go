package models

type Zone struct {
	Id string `json:"id" bson:"id" validate:"required"`
}
