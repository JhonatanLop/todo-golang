package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ListUser); err != nil {
		http.Error(w, "Filed to encode users", http.StatusInternalServerError)
	}
}

func StartServer() {
	http.HandleFunc("/users", getUsers)
	http.ListenAndServe(":8080", nil)
}
