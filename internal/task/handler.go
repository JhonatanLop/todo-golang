package task

import "net/http"

func TaskHandler(w http.ResponseWriter, r *http.Request) {
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
