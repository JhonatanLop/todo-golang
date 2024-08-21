package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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
			return nil
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Task not found"))
	return nil
}

func postTask(w http.ResponseWriter, r *http.Request) error {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		return err
	}
	if newTask.Id <= ListTask[len(ListTask)-1].Id {
		newTask.Id = ListTask[len(ListTask)-1].Id + 1
		ListTask = append(ListTask, newTask)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task added sucessfully, but the id was changed because it was invalid"))
		fmt.Println(newTask)
		return nil
	}

	ListTask = append(ListTask, newTask)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task added successfully"))
	fmt.Println(newTask)
	return nil
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Service error! ::Failed to decode the task", http.StatusBadRequest)
	}
	for i, task := range ListTask {
		if task.Id == updatedTask.Id {
			ListTask[i] = updatedTask
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Task was updated successfully"))
			return
		}
	}
	http.Error(w, "Cannot find the task", http.StatusNotFound)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	target, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid param", http.StatusBadRequest)

	}
	for i, task := range ListTask {
		if task.Id == target {
			ListTask = append(ListTask[:i], ListTask[i+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Task removed successfully"))
			return
		}
	}
	http.Error(w, "Cannot find the task", http.StatusNotFound)
}
