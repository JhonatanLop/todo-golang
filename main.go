package main

import (
	"fmt"
	"time"
)

func main() {
	// criando um Usuário
	newUser := CreateUser("Jhow", "jhow@email.com", "senha123")

	// criando uma task
	newTask := CreateTask(1, "Title", "Description", newUser, time.Now(), time.Now(), time.Now(), 10, "red")
	otherTask := CreateTask(1, "Title", "Description", newUser, time.Now(), time.Now(), time.Now(), 10, "red")
	anotherTask := CreateTask(1, "Title", "Description", newUser, time.Now(), time.Now(), time.Now(), 10, "red")

	ListTask = append(ListTask, newTask)
	ListTask = append(ListTask, otherTask)
	ListTask = append(ListTask, anotherTask)

	fmt.Println(ListTask)

	// UpdateTask()
}
