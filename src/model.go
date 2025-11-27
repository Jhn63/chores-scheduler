package main

import "time"

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"Title"`
	Description string `json:"description"`
	Repetable   bool   `json:"repetable"`
}

type TaskHandler struct {
	LastDate time.Time `json:"last-date"`
	TodoTask []Task    `json:"todo-tasks"`
	AllTasks []Task    `json:"all-tasks"`
}
