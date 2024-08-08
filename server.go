package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ListUser); err != nil {
		http.Error(w, "Filed to encode users", http.StatusInternalServerError)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for _, user := range ListUser {
		if user.Id == id {
			if err := json.NewEncoder(w).Encode(user); err != nil {
				http.Error(w, "Failed to encode user", http.StatusInternalServerError)
			}
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func StartServer() {
	http.HandleFunc("/users", getAllUsers)
	http.HandleFunc("/user", getUser)
	http.ListenAndServe(":8080", nil)
}
