package main

import (
	"time"
)

func main() {
	// criando um Usuário
	newUser := CreateUser("Jhow", "jhow@email.com", "senha123")

	// criando uma task
	newTask := CreateTask(1, "Title", "Description", newUser, time.Now(), time.Now(), time.Now(), 10, "red")
	otherTask := CreateTask(2, "Title", "Description", newUser, time.Now(), time.Now(), time.Now(), 10, "red")
	anotherTask := CreateTask(3, "Title", "Description", newUser, time.Now(), time.Now(), time.Now(), 10, "red")

	ListTask = append(ListTask, newTask)
	ListTask = append(ListTask, otherTask)
	ListTask = append(ListTask, anotherTask)

	ShowTask()

	otherUser := CreateUser("Sebastian", "sebastian@email.com", "senha123")
	newAnotherTask := CreateTask(1, "Another title", "My description", otherUser, time.Now(), time.Now(), time.Now(), 10, "red")
	UpdateTask(newAnotherTask)

	println("\n----\n")

	ShowTask()
}
