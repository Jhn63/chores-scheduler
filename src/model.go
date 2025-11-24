package main

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"Title"`
	Description string `json:"description"`
	Repetable   bool   `json:"repetable"`
}

type TaskHandler struct {
	TodoTask []Task `json:"todo-tasks"`
	AllTasks []Task `json:"all-tasks"`
}
