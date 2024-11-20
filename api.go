package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type message struct {
	Text string `json:"text"`
}

func main() {

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := message{
			Text: "Front end connected to Go Api Server",	 
		}
		json.NewEncoder(w).Encode(response)
	})

	// to load static front end page in our server
	http.Handle("/", http.FileServer(http.Dir("./public")))

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
