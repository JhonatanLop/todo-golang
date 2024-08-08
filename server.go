package main

import (
	"net/http"
)

func StartServer() {
	http.HandleFunc("/user", UserHandler)
	http.ListenAndServe(":8080", nil)
}
