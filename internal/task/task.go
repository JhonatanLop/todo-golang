package task

import (
	"time"
)

var ListTask []Task

type Task struct {
	Id            int
	Title         string
	Description   string
	DueDate       time.Time
	CreateDate    time.Time
	CompletedDate time.Time
	Difficulty    uint8
}

func CreateTask(
	id int,
	title string,
	description string,
	dueDate time.Time,
	createDate time.Time,
	completedDate time.Time,
	difficulty uint8,
) Task {
	return Task{
		Id:            id,
		Title:         title,
		Description:   description,
		DueDate:       dueDate,
		CreateDate:    createDate,
		CompletedDate: completedDate,
		Difficulty:    difficulty,
	}
}
