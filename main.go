package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type RequestBody struct {
	Name   string `json:"name"`
	Taijyu string `json:"taijyu"`
}

func recordHandler(w http.ResponseWriter, r *http.Request) {
	secretToken := os.Getenv("SECRET_TOKEN")
	authHeader := r.Header.Get("Authorization")

	if authHeader != "Bearer "+secretToken {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
		return
	}

	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
		return
	}

	log.Printf("受け取った: %s, %s", body.Name, body.Taijyu)
	json.NewEncoder(w).Encode(map[string]string{"message": "受け取ったっちゃ！"})
}

func main() {
	http.HandleFunc("/record", recordHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s...", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/record", recordHandler)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
