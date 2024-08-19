package task

var ListTask []Task

type Task struct {
	Id            int
	Title         string
	Description   string
	DueDate       string
	CreateDate    string
	CompletedDate string
	Difficulty    uint8
}

func CreateTask(
	id int,
	title string,
	description string,
	dueDate string,
	createDate string,
	completedDate string,
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
