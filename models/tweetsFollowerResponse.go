package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TweetsFollowerResponse struct {
	ID             primitive.ObjectID `bson: "_id" json: "_id,omitempty"`
	UserId         string             `bson: "userid", json:"userid,omitempty"`
	UserRelationId string             `bson: "userrelationid", json:"userrelationid,omitempty"`
	Tweet          struct {
		Message string    `bson: "message", json:"message,omitempty"`
		Date    time.Time `bson: "date", json:"date,omitempty"`
		ID      string    `bson: "_id", json:"_id,omitempty"`
	}
}
