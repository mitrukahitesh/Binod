package category

import (
	"context"
	"ecell-server/db"
	"ecell-server/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Category model.Category `json:"category"`
}

func GetAll(w http.ResponseWriter, r *http.Request, cat string) {
	client, err := db.GetMongoClient()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	collection := client.Database(db.DB).Collection(db.CATEGORY)
	var categories model.Category
	filter := bson.D{primitive.E{Key: "category", Value: cat}}
	res := collection.FindOne(context.TODO(), filter)
	err = res.Decode(&categories)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(Response{
		Category: categories,
	})
}
