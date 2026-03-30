package main

import "time"

type Task struct {
	Id                 int
	Title              string
	Description        string
	DegreeOfDifficulty int
	DegreeOfImportance int
	Deadline           *time.Time
	LastQueuedAt       time.Time
	Repeatable         bool
	Active             bool
}

type TaskRead struct {
	Id                 int        `json:"id"`
	Title              string     `json:"title"`
	Description        string     `json:"description"`
	DegreeOfDifficulty int        `json:"degree_of_difficulty"`
	Deadline           *time.Time `json:"deadline"`
}

type TaskCreate struct {
	Title              string     `json:"title"`
	Description        string     `json:"description"`
	DegreeOfDifficulty int        `json:"degree_of_difficulty"`
	DegreeOfImportance int        `json:"degree_of_importance"`
	Deadline           *time.Time `json:"deadline"`
	Repeatable         bool       `json:"repeatable"`
}

type TaskUpdate struct {
	Title              *string    `json:"title"`
	Description        *string    `json:"description"`
	DegreeOfDifficulty *int       `json:"degree_of_difficulty"`
	DegreeOfImportance *int       `json:"degree_of_importance"`
	Deadline           *time.Time `json:"deadline"`
	Repeatable         *bool      `json:"repeatable"`
}
