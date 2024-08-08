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

// http handler

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		if len(r.URL.Query()) == 0 {
			getAllUsers(w, r)
		} else {
			getUser(w, r)
		}
	case "POST":
		postUser(w, r)
	case "PUT":
		putUser(w, r)
	case "DELETE":
		deleteUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// http methods

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(ListUser); err != nil {
		http.Error(w, "Filed to encode users", http.StatusInternalServerError)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
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

func putUser(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println(err)
		return
	}
}

func postUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Failed to encode user", http.StatusBadRequest)
		fmt.Println(err)
	}
	fmt.Println(newUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	idTg, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Filed to identify user", http.StatusBadRequest)
		fmt.Println(err)
	}
	for i, user := range ListUser {
		if user.Id == idTg {
			ListUser = append(ListUser[:i], ListUser[:i+1]...)
		}
	}
}
