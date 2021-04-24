package routes

import (
	"ecell-server/routes/banner"
	"ecell-server/routes/category"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	query := r.URL.Query()
	resource := query.Get("resource")
	cat := query.Get("category")
	switch resource {
	case "banner":
		banner.GetAll(w, r)
	case "category":
		category.GetAll(w, r, cat)
	}
}
