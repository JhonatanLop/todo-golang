package main

import (
	"time"
)

func main() {
	// criando um Usuário
	newUser := CreateUser(1, "Jhow", "jhow@email.com", "senha123")
	ListUser = append(ListUser, newUser)

	// criando uma task
	newTask := CreateTask(1, "Title", "Description", newUser, time.Now(), time.Now(), time.Now(), 10, "red")
	otherTask := CreateTask(2, "Title", "Description", newUser, time.Now(), time.Now(), time.Now(), 10, "red")
	anotherTask := CreateTask(3, "Title", "Description", newUser, time.Now(), time.Now(), time.Now(), 10, "red")

	ListTask = append(ListTask, newTask)
	ListTask = append(ListTask, otherTask)
	ListTask = append(ListTask, anotherTask)

	// ShowTask()

	otherUser := CreateUser(2, "Sebastian", "sebastian@email.com", "senha123")
	ListUser = append(ListUser, otherUser)

	newAnotherTask := CreateTask(1, "Another title", "My description", otherUser, time.Now(), time.Now(), time.Now(), 10, "red")
	UpdateTask(newAnotherTask)

	// println("\n----\n")

	// ShowTask()

	anotherUser := CreateUser(3, "Kleiton", "kleiton@email.com", "senha123")
	ListUser = append(ListUser, anotherUser)

	ShowUsers()
	println("\n----\n")
	newAnotherUser := CreateUser(3, "Teseu", "teseu@email.com", "senha123")
	UpdateUser(newAnotherUser)
	ShowUsers()
}
