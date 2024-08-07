package main

import (
	"fmt"
	"time"
)

var ListTask []Task

type Task struct {
	Id            int
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
	id int,
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
		Id:            id,
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

func ShowTask() {
	for i := range ListTask {
		fmt.Println(ListTask[i])
	}
}

func UpdateTask(task Task) {
	for i := range ListTask {
		if ListTask[i].Id == task.Id {
			ListTask[i] = task
		}
	}
}

func DeleteTask(id int) {
	for i := range ListTask {
		if ListTask[i].Id == id {
			ListTask = append(ListTask[:i], ListTask[i+1:]...)
		}
	}
}
