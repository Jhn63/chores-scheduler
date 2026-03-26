package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func createTaskService(task TaskCreate) (Task, error) {
	stmt, err := DB.Prepare("INSERT INTO tasks (title, description, degree_of_difficulty, deadline, repeatable, active) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return Task{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(task.Title, task.Description, task.DegreeOfDifficulty, task.Deadline, task.Repeatable, true)
	if err != nil {
		return Task{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Task{}, err
	}

	createdTask := Task{
		Id:                 int(id),
		Title:              task.Title,
		Description:        task.Description,
		DegreeOfDifficulty: task.DegreeOfDifficulty,
		Deadline:           task.Deadline,
		Repeatable:         task.Repeatable,
		Active:             true,
	}

	return createdTask, nil
}

func updateTaskService(id int, task TaskUpdate) (Task, error) {
	var updatedTask Task

	var setClauses []string
	var args []any

	if task.Title != nil {
		setClauses = append(setClauses, "title = ?")
		args = append(args, *task.Title)
	}
	if task.Description != nil {
		setClauses = append(setClauses, "description = ?")
		args = append(args, *task.Description)
	}
	if task.DegreeOfDifficulty != nil {
		setClauses = append(setClauses, "degree_of_difficulty = ?")
		args = append(args, *task.DegreeOfDifficulty)
	}
	if task.Deadline != nil {
		setClauses = append(setClauses, "deadline = ?")
		args = append(args, *task.Deadline)
	}
	if task.Repeatable != nil {
		setClauses = append(setClauses, "repeatable = ?")
		args = append(args, *task.Repeatable)
	}

	if len(setClauses) == 0 {
		return updatedTask, errors.New("no fields provided to update")
	}

	query := "UPDATE tasks SET " + strings.Join(setClauses, ", ") + " WHERE id = ?"

	args = append(args, id)

	_, err := DB.Exec(query, args...)
	if err != nil {
		return updatedTask, err
	}

	row := DB.QueryRow(`
		SELECT id, title, description, degree_of_difficulty, deadline, repeatable, active 
		FROM tasks WHERE id = ?`, id,
	)

	err = row.Scan(
		&updatedTask.Id,
		&updatedTask.Title,
		&updatedTask.Description,
		&updatedTask.DegreeOfDifficulty,
		&updatedTask.Deadline,
		&updatedTask.Repeatable,
		&updatedTask.Active,
	)

	return updatedTask, err
}

func deleteTaskService(id int) (Task, error) {
	var task Task
	err := DB.QueryRow("SELECT id, title, description, degree_of_difficulty, deadline, repeatable, active FROM tasks WHERE id = ?", id).Scan(
		&task.Id, &task.Title, &task.Description, &task.DegreeOfDifficulty, &task.Deadline, &task.Repeatable, &task.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return Task{}, fmt.Errorf("task with id %d not found", id)
		}
		return Task{}, err
	}

	// Delete the task
	_, err = DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func getTasksService(limit int) ([]TaskRead, error) {
	rows, err := DB.Query(`SELECT id, title, description, degree_of_difficulty, deadline FROM tasks WHERE active = true ORDER BY last_queued_at ASC LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []TaskRead
	for rows.Next() {
		var task TaskRead
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.DegreeOfDifficulty, &task.Deadline)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func getTaskByIDService(id int) (TaskRead, error) {
	var task TaskRead
	err := DB.QueryRow("SELECT id, title, description, degree_of_difficulty, deadline FROM tasks WHERE id = ?", id).Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.DegreeOfDifficulty,
		&task.Deadline,
	)
	if err != nil {
		return TaskRead{}, err
	}
	return task, nil
}

func getAllTasksService() ([]TaskRead, error) {
	rows, err := DB.Query("SELECT id, title, description, degree_of_difficulty, deadline FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []TaskRead
	for rows.Next() {
		var task TaskRead
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.DegreeOfDifficulty, &task.Deadline)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func setDoneTaskService(id int) (TaskRead, error) {
	_, err := DB.Exec("UPDATE tasks SET active = CASE WHEN repeatable = TRUE THEN TRUE ELSE FALSE END, last_queued_at = CASE WHEN repeatable = TRUE THEN CURRENT_TIMESTAMP ELSE last_queued_at END WHERE id = ?", id)
	if err != nil {
		return TaskRead{}, err
	}

	return getTaskByIDService(id)
}
