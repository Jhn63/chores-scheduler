package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func startServer() {
	router := http.NewServeMux()
	router.HandleFunc("GET /hello", home)
	router.HandleFunc("GET /tasks", getTasks)
	router.HandleFunc("GET /tasks/{id}", getTaskByID)
	router.HandleFunc("GET /tasks/all", getAllTasks)
	router.HandleFunc("POST /tasks", createTask)
	router.HandleFunc("PUT /tasks/{id}", updateTask)
	router.HandleFunc("POST /tasks/{id}/done", setDoneTask)
	router.HandleFunc("DELETE /tasks/{id}", deleteTask)

	fmt.Println("Server is running on port 8000...")
	handler := corsMiddleware(router)
	if err := http.ListenAndServe(":8000", handler); err != nil {
		panic(err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests, send 200 OK for any OPTIONS method
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Lazy Tasks Manager!")
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var input TaskCreate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	task, err := createTaskService(input)
	if err != nil {
		http.Error(w, "Fail Creating Task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")
	taskID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid task ID. Must be a number.", http.StatusBadRequest)
		return
	}

	var input TaskUpdate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	task, err := updateTaskService(taskID, input)
	if err != nil {
		http.Error(w, "Fail Updating Task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	taskID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid task ID. Must be a number.", http.StatusBadRequest)
		return
	}

	task, err := deleteTaskService(taskID)
	if err != nil {
		http.Error(w, "Fail Deleting Task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func getTasks(w http.ResponseWriter, r *http.Request) {

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10 // Default limit
	}

	tasks, err := getTasksService(limit)
	if err != nil {
		http.Error(w, "Fail Getting Tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func getTaskByID(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	taskID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid task ID. Must be a number.", http.StatusBadRequest)
		return
	}

	task, err := getTaskByIDService(taskID)
	if err != nil {
		http.Error(w, "Fail Getting Task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := getAllTasksService()
	if err != nil {
		http.Error(w, "Fail Getting Tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func setDoneTask(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	taskID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid task ID. Must be a number.", http.StatusBadRequest)
		return
	}

	task, err := setDoneTaskService(taskID)
	if err != nil {
		http.Error(w, "Fail Setting Task as Done", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
