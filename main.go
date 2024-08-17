package main

import (
	"log"
	"net/http"

	task "github.com/JhonatanLop/todo-golang/internal/task"
	//  user "github.com/JhonatanLop/todo-golang/internal/user"
	"time"
)

func main() {
	StartServer()
	initSeeds()
}

func StartServer() error {
	// definindo handler
	// http.HandleFunc("/user", UserHandler)
	http.HandleFunc("/task", task.TaskHandler)

	// iniciando o servidor
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server: %v", err)
		return err
	}
	return nil
}

func initSeeds() {
	// criando um Usuário
	// newUser := CreateUser(1, "Jhow", "jhow@email.com", "senha123")
	// ListUser = append(ListUser, newUser)

	var mytime time.Time
	var difficulty uint8
	// criando uma task
	newTask := task.CreateTask(1, "Title", "Description", mytime, time.Now(), mytime, difficulty)
	otherTask := task.CreateTask(2, "Title", "Description", mytime, time.Now(), mytime, difficulty)
	anotherTask := task.CreateTask(3, "Title", "Description", mytime, time.Now(), mytime, difficulty)
	newAnotherTask := task.CreateTask(4, "Another title", "My description", mytime, time.Now(), mytime, difficulty)

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
