package main

type User struct {
	Name  string
	Email string
	Pw    string
}

func CreateUser(name string, email string, pw string) User {
	return User{
		Name:  name,
		Email: email,
		Pw:    pw,
	}
}
