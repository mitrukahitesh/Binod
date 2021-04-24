package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Banner struct {
	Id   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Link string             `json:"link" bson:"link"`
}
