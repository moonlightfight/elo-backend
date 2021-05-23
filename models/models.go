package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

type Player struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Slug       string             `json:"slug,omitempty" bson:"slug,omitempty"`
	Username   string             `json:"username,omitempty" bson:"username,omitempty"`
	Country    string             `json:"country,omitempty" bson:"country,omitempty"`
	Ranking    int32              `json:"ranking,omitempty" bson:"ranking,omitempty"`
	Points     int32              `json:"points,omitempty" bson:"points,omitempty"`
	Controller string             `json:"controller,omitempty" bson:"controller,omitempty"`
	RealName   string             `json:"realName,omitempty" bson:"realName,omitempty"`
	Twitter    string             `json:"twitter,omitempty" bson:"twitter,omitempty`
	Twitch     string             `json:"twitch,omitempty" bson:"twitch,omitempty"`
}

type LoginData struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}