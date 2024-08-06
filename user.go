package main

import "fmt"

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
