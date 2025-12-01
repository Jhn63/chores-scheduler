package main

import (
	"fmt"
	"os"
)

func printTask(task Task) {
	if !task.Repetable {
		fmt.Print("* ")
	}
	fmt.Printf("%d - %s:\n", task.Id, task.Title)
	fmt.Printf("%s\n\n", task.Description)
}

func main() {
	th := TaskHandler{}
	th.Init()

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command:")
		fmt.Println("\t- list-tasks:\n\t\tLists all available tasks.\n")
		fmt.Println("\t- get-tasks:\n\t\tLists today's tasks.\n")
		fmt.Println("\t- set-done <id>:\n\t\tMarks the task with the given ID as done.\n")
		fmt.Println("\t- delete-task <id>:\n\t\tDeletes the task with the given ID.\n")
		fmt.Println("\t- new-task <id> <title> <description> <repetable>:\n\t\tCreates a new task with the provided details.\n")
		return
	}

	switch os.Args[1] {
	case "list-tasks":
		all := append(th.AllTasks, th.TodoTask...)
		if len(all) == 0 {
			fmt.Println("No tasks available.")

		} else {

			for _, task := range all {
				printTask(task)
			}
		}

	case "get-tasks":
		if len(th.TodoTask) == 0 {
			fmt.Println("No more tasks for today.")

		} else {
			fmt.Println("Today's tasks:")
			for _, task := range th.TodoTask {
				printTask(task)
			}
		}

	case "set-done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the task ID to set as done.")
			return
		}
		var id int
		if _, err := fmt.Sscanf(os.Args[2], "%d", &id); err != nil {
			fmt.Println("Invalid task ID.")
			return
		}

		th.SetTaskDone(id)

	case "delete-task":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the task ID to delete.")
			return
		}
		var id int
		if _, err := fmt.Sscanf(os.Args[2], "%d", &id); err != nil {
			fmt.Println("Invalid task ID.")
			return
		}

		th.DeleteTask(id)

	case "new-task":
		if len(os.Args) < 5 {
			fmt.Println("Please provide the task ID, title, description, and repetable status.")
			return
		}
		var id int
		var title, description string
		var repetable bool
		if _, err := fmt.Sscanf(os.Args[2], "%d", &id); err != nil {
			fmt.Println("Invalid task ID.")
			return
		}
		title = os.Args[3]
		description = os.Args[4]
		if _, err := fmt.Sscanf(os.Args[5], "%t", &repetable); err != nil {
			fmt.Println("Invalid repetable status. Use true or false.")
			return
		}

		th.NewTask(id, title, description, repetable)

	default:
		fmt.Println("Unknown command.")
		return

	}
	th.Save()
}
