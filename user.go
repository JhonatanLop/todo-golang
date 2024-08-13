package main

import (
	"encoding/json"
	"errors"
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
		if err := handleGetRequest(w, r); err != nil {
			errors.New(err.Error())
		}
	case "POST":
		if err := postUser(w, r); err != nil {
			errors.New(err.Error())
		}
	case "PUT":
		if err := putUser(w, r); err != nil {
			errors.New(err.Error())
		}
	case "DELETE":
		if err := deleteUser(w, r); err != nil {
			errors.New(err.Error())
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) error {
	params := r.URL.Query()
	if params.Has("id") {
		user, err := getUser(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "Invalid parameter", http.StatusBadRequest)
			return err
		}
	} else if params.Has("email") {
		user, err := getUserByEmail(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "Invalid parameter", http.StatusBadRequest)
			return err
		}
	} else {
		getAllUsers(w, r)
	}
	return nil
}

func getAllUsers(w http.ResponseWriter, r *http.Request) ([]User, error) {
	if err := json.NewEncoder(w).Encode(ListUser); err != nil {
		http.Error(w, "Filed to encode users", http.StatusInternalServerError)
		return []User{}, errors.New(err.Error())
	}
	return ListUser, nil
}

func getUser(w http.ResponseWriter, r *http.Request) (User, error) {
	idStr := r.URL.Query().Get("id")

	//
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return User{}, err
	}

	for _, user := range ListUser {
		if user.Id == id {
			if err := json.NewEncoder(w).Encode(user); err != nil {
				return User{}, err
			}
			return user, nil
		}
	}
	return User{}, errors.New("User not found")
}

func getUserByEmail(w http.ResponseWriter, r *http.Request) (User, error) {
	emailTarget := r.URL.Query().Get("email")
	for _, user := range ListUser {
		if user.Email == emailTarget {
			if err := json.NewEncoder(w).Encode(user); err != nil {
				http.Error(w, "Failed to encode user", http.StatusInternalServerError)
				return User{}, err
			}
			return user, nil
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
	return User{}, errors.New("User not found")
}

func putUser(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	var newUser User
	// validação de requisição
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return err
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
		return err
	}
	return nil
}

func postUser(w http.ResponseWriter, r *http.Request) error {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Failed to encode user", http.StatusBadRequest)
		return err
	}
	ListUser = append(ListUser, newUser)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User added successfully"))
	return nil
}

func deleteUser(w http.ResponseWriter, r *http.Request) error {
	idTg, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid param", http.StatusBadRequest)
		return err
	}
	for i, user := range ListUser {
		if user.Id == idTg {
			ListUser = append(ListUser[:i], ListUser[:i+1]...)
		}
	}
	return nil
}
