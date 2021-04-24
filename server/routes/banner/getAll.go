package banner

import (
	"context"
	"ecell-server/db"
	"ecell-server/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type Response struct {
	Banners []model.Banner `json:"banners"`
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	client, err := db.GetMongoClient()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	collection := client.Database(db.DB).Collection(db.BANNER)
	var banners []model.Banner
	filter := bson.D{{}}
	res, err := collection.Find(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for res.Next(context.TODO()) {
		var b model.Banner
		err := res.Decode(&b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		banners = append(banners, b)
	}
	json.NewEncoder(w).Encode(Response{
		Banners: banners,
	})
}
