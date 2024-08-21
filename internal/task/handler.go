package task

import "net/http"

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	switch r.Method {
	case "GET":
		taskGetHandler(w, r)
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
