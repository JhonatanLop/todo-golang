package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var ListUser []User

type User struct {
	Id    int
	Name  string
	Email string
	Pw    string
}

func CreateUser(id int, name string, email string, pw string) User {
	return User{
		Id:    id,
		Name:  name,
		Email: email,
		Pw:    pw,
	}
}

func ShowUsers() {
	for i := range ListUser {
		fmt.Println(ListUser[i])
	}
}

func DeleteUser(user User) {
	for i := range ListUser {
		if ListUser[i].Id == user.Id {
			ListUser = append(ListUser[:i], ListUser[i+1:]...)
		}
	}
}

func UpdateUser(user User) {
	for i := range ListUser {
		if ListUser[i].Id == user.Id {
			ListUser[i] = user
		}
	}
}

// http methods

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ListUser); err != nil {
		http.Error(w, "Filed to encode users", http.StatusInternalServerError)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
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

func PutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser User
	// validação de requisição
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// adicionando o novo usuário à lista
	for i := range ListUser {
		if ListUser[i].Id == newUser.Id {
			ListUser[i] = newUser
			break
		}
	}

	// se usuário não existe, adiciona na lista
	ListUser = append(ListUser, newUser)

	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}
