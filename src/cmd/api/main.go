package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message" bson:"message"`
}

func main() {
	fmt.Println("Listening on: http://localhost:8000")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Response{Message: "pong"})
	})
	
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Failure.")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<div><h1>Hello K8s</h1></div>"))
}
