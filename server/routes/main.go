package routes

import (
	"ecell-server/routes/banner"
	"ecell-server/routes/category"
	"ecell-server/routes/email"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	query := r.URL.Query()
	resource := query.Get("resource")
	cat := query.Get("category")
	mail := query.Get("email")
	name := query.Get("name")
	switch resource {
	case "banner":
		banner.GetAll(w, r)
	case "category":
		category.GetAll(w, r, cat)
	case "sendmail":
		email.SendMail(w, r, mail, name)
	}
}
