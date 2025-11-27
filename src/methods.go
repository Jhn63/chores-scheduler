package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const MAX_TODO_TASKS = 5

func (th *TaskHandler) Init() {
	th.Load("saves.json")
	defer th.Save()

	currentDate := time.Now().Truncate(24 * time.Hour)
	if th.LastDate.IsZero() || th.LastDate.Before(currentDate) {
		th.LastDate = currentDate
		th.ClearTasks()
		th.GetTasks(MAX_TODO_TASKS)
	}

	if len(th.TodoTask) < MAX_TODO_TASKS {
		th.GetTasks(MAX_TODO_TASKS - len(th.TodoTask))
	}
}

func (th *TaskHandler) Load(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		th.Save()
		return
	}

	if err := json.Unmarshal(data, &th); err != nil {
		log.Fatal("Couldn't unmarshal json")
	}
}

func (th *TaskHandler) Save() {
	json, err := json.Marshal(&th)
	if err != nil {
		fmt.Println("Couldn't marshal json")
	}

	//look permissions
	if erro := os.WriteFile("saves.json", json, 0644); erro != nil {
		fmt.Println("Couldn't save json")
	}
}

func (th *TaskHandler) GetTasks(num int) {
	th.TodoTask = append(th.TodoTask, th.AllTasks[:num]...) //inserting in todo tasks
	th.AllTasks = th.AllTasks[num:]                         //removing from all tasks
}

func (th *TaskHandler) ClearTasks() {
	th.AllTasks = append(th.AllTasks, th.TodoTask...)
	th.TodoTask = []Task{}
}

func (th *TaskHandler) NewTask(id int, title string, dscrptn string, reapt bool) {
	task := Task{
		Id:          id,
		Title:       title,
		Description: dscrptn,
		Repetable:   reapt,
	}
	th.AllTasks = append(th.AllTasks, task)
}

func (th *TaskHandler) SetTaskDone(id int) {

	for i, task := range th.TodoTask {
		if task.Id == id {

			th.TodoTask = append(th.TodoTask[:i], th.TodoTask[i+1:]...) //removing from todo tasks

			if task.Repetable {
				th.AllTasks = append(th.AllTasks, task) //adding back to all tasks
			}
			break
		}
	}
}

func (th *TaskHandler) DeleteTask(id int) {

	for i, task := range th.AllTasks {
		if task.Id == id {
			th.AllTasks = append(th.AllTasks[:i], th.AllTasks[i+1:]...)
			break
		}
	}
	for i, task := range th.TodoTask {
		if task.Id == id {
			th.TodoTask = append(th.TodoTask[:i], th.TodoTask[i+1:]...)
			break
		}
	}
}
