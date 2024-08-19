package main

import (
	"log"
	"net/http"

	task "github.com/JhonatanLop/todo-golang/internal/task"
	//  user "github.com/JhonatanLop/todo-golang/internal/user"
	"time"
)

func main() {
	initSeeds()
	StartServer()
}

func StartServer() {
	// definindo handler
	// http.HandleFunc("/user", UserHandler)
	http.HandleFunc("/task", task.TaskHandler)

	// iniciando o servidor
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server: %v", err)
	}
}

func initSeeds() {
	// criando um Usuário
	// newUser := CreateUser(1, "Jhow", "jhow@email.com", "senha123")
	// ListUser = append(ListUser, newUser)

	// criando uma task
	var thisTime = time.Now().String()
	newTask := task.CreateTask(1, "Title", "Description", thisTime, thisTime, thisTime, 5)
	otherTask := task.CreateTask(2, "Title", "Description", thisTime, thisTime, thisTime, 5)
	anotherTask := task.CreateTask(3, "Title", "Description", thisTime, thisTime, thisTime, 5)
	newAnotherTask := task.CreateTask(4, "Another title", "My description", thisTime, thisTime, thisTime, 5)

	// otherUser := CreateUser(2, "Sebastian", "sebastian@email.com", "senha123")
	// anotherUser := CreateUser(3, "Kleiton", "kleiton@email.com", "senha123")
	// newAnotherUser := CreateUser(4, "Teseu", "teseu@email.com", "senha123")

	task.ListTask = append(task.ListTask, newTask)
	task.ListTask = append(task.ListTask, otherTask)
	task.ListTask = append(task.ListTask, anotherTask)
	task.ListTask = append(task.ListTask, newAnotherTask)

	// ListUser = append(ListUser, otherUser)
	// ListUser = append(ListUser, anotherUser)
	// ListUser = append(ListUser, newAnotherUser)
}
