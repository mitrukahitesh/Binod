package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	Name  string `json:"name" bson:"name"`
	Offer string `json:"offer" bson:"offer"`
	Img   string `json:"img" bson:"img"`
}

type Category struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type  string             `json:"category" bson:"category"`
	Items []Item             `json:"items" bson:"items"`
}
