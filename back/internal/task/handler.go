package task

import (
	"fmt"
	"net/http"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	switch r.Method {
	case "GET":
		if err := taskGetHandler(w, r); err != nil {
			fmt.Println(err)
		}
	case "POST":
		postTask(w, r)
	case "PUT":
		updateTask(w, r)
	case "DELETE":
		deleteTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func taskGetHandler(w http.ResponseWriter, r *http.Request) error {
	var errHandler error
	param := r.URL.Query()
	if param.Has("id") {
		errHandler = getTaskById(w, r)
	} else if param.Has("desc") {
		errHandler = getSortedDescTask(w, r)
	} else {
		errHandler = getAllTask(w, r)
	}
	return errHandler
}
