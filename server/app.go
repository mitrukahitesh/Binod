package main

import (
	"ecell-server/db"
	"ecell-server/routes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//CONNECT MONGO CLUSTER
	_, err := db.GetMongoClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	//STARTING SERVER
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.Handler)
	fmt.Println("[STARTED] Server started...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
}
