package main

import (
	"time"
)

func main() {
	// criando um Usuário
	newUser := CreateUser(1, "Jhow", "jhow@email.com", "senha123")
	ListUser = append(ListUser, newUser)

	var mytime *time.Time
	var difficulty *uint8
	// criando uma task
	newTask := CreateTask(1, "Title", "Description", mytime, time.Now(), mytime, difficulty)
	otherTask := CreateTask(2, "Title", "Description", mytime, time.Now(), mytime, difficulty)
	anotherTask := CreateTask(3, "Title", "Description", mytime, time.Now(), mytime, difficulty)

	ListTask = append(ListTask, newTask)
	ListTask = append(ListTask, otherTask)
	ListTask = append(ListTask, anotherTask)

	otherUser := CreateUser(2, "Sebastian", "sebastian@email.com", "senha123")
	ListUser = append(ListUser, otherUser)

	newAnotherTask := CreateTask(1, "Another title", "My description", mytime, time.Now(), mytime, difficulty)
	ListTask = append(ListTask, newAnotherTask)
	// UpdateTask(newAnotherTask)

	anotherUser := CreateUser(3, "Kleiton", "kleiton@email.com", "senha123")
	ListUser = append(ListUser, anotherUser)

	newAnotherUser := CreateUser(4, "Teseu", "teseu@email.com", "senha123")
	ListUser = append(ListUser, newAnotherUser)

	StartServer()
}
