package main

import (
	"time"
)

type Task struct {
	Index         int
	Title         string
	Description   string
	User          User
	DueDate       time.Time
	CreateDate    time.Time
	CompletedDate time.Time
	Difficulty    uint8
	Color         string
}

func CreateTask(
	index int,
	title string,
	description string,
	user User,
	dueDate time.Time,
	createDate time.Time,
	completedDate time.Time,
	difficulty uint8,
	color string,
) Task {
	return Task{
		Index:         index,
		Title:         title,
		Description:   description,
		User:          user,
		DueDate:       dueDate,
		CreateDate:    createDate,
		CompletedDate: completedDate,
		Difficulty:    difficulty,
		Color:         color,
	}
}
