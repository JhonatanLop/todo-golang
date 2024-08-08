package main

import (
	"net/http"
)

func StartServer() {
	http.HandleFunc("/users", GetAllUsers)
	http.HandleFunc("/user", GetUser)
	http.ListenAndServe(":8080", nil)
}
