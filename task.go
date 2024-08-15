package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

var ListTask []Task

type Task struct {
	Id            int
	Title         string
	Description   string
	DueDate       *time.Time
	CreateDate    time.Time
	CompletedDate *time.Time
	Difficulty    *uint8
}

func CreateTask(
	id int,
	title string,
	description string,
	dueDate *time.Time,
	createDate time.Time,
	completedDate *time.Time,
	difficulty *uint8,
) Task {
	return Task{
		Id:            id,
		Title:         title,
		Description:   description,
		DueDate:       dueDate,
		CreateDate:    createDate,
		CompletedDate: completedDate,
		Difficulty:    difficulty,
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	switch r.Method {
	case "GET":
		if err := taskGetHandler(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case "POST":
		if err := postTask(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case "PUT":
		if err := putTask(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case "DELETE":
		if err := deleteTask(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func taskGetHandler(w http.ResponseWriter, r *http.Request) error {
	param := r.URL.Query()
	if param.Has("id") {
		return getTaskById(w, r)
	} else {
		return getAllTask(w, r)
	}
}

func getAllTask(w http.ResponseWriter, r *http.Request) error {
	if err := json.NewEncoder(w).Encode(ListTask); err != nil {
		return err
	}
	return nil
}

func getTaskById(w http.ResponseWriter, r *http.Request) error {
	taskId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return err
	}

	for _, task := range ListTask {
		if task.Id == taskId {
			if err := json.NewEncoder(w).Encode(task); err != nil {
				return err
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Task not found"))
	return nil
}

func postTask(w http.ResponseWriter, r *http.Request) error {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(newTask); err != nil {
		return err
	}
	ListTask = append(ListTask, newTask)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task added successfully"))
	return nil
}

func putTask(w http.ResponseWriter, r *http.Request) error {
	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(updatedTask); err != nil {
		return err
	}
	for _, task := range ListTask {
		if task.Id == updatedTask.Id {
			task = updatedTask
			return nil
		}
	}
	http.Error(w, "Cannot find the task", http.StatusNotFound)
	return nil
}

func deleteTask(w http.ResponseWriter, r *http.Request) error {
	var target Task
	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		return err
	}
	for i, task := range ListTask {
		if task.Id == target.Id {
			ListTask = append(ListTask[:i], ListTask[i+1:]...)
			return nil
		}
	}
	http.Error(w, "Cannot find the task", http.StatusNotFound)
	return nil
}
