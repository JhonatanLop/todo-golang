package main

import (
	"log"
	"net/http"
)

func StartServer() {
	// definindo handler
	http.HandleFunc("/user", UserHandler)

	// iniciando o servidor
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server: %v", err)
	}
}
