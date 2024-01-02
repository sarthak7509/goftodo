package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID             uuid.UUID `gorm:"PrimaryKey"`
	Title          string
	SubTitle       string
	EstimateSprint string
	IsDone         bool
	Priority       int
	UserId         uuid.UUID
	User           User `gorm:"constraint:OnDelete:CASCADE;"`
}

type TodoCreatePayload struct {
	Title          string `json:"title"`
	SubTitle       string `json:"sub_title"`
	EstimateSprint string `json:"estimate_sprint"`
	IsDone         bool   `json:"is_done"`
	Priority       int    `json:"priority"`
}
type TodoGetPayload struct {
	Id             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	SubTitle       string    `json:"sub_title"`
	EstimateSprint string    `json:"estimate_sprint"`
	IsDone         bool      `json:"is_done"`
	Priority       int       `json:"priority"`
}

func TodoGetResponse(todo Todo) TodoGetPayload {
	return TodoGetPayload{
		Title:          todo.Title,
		SubTitle:       todo.SubTitle,
		EstimateSprint: todo.EstimateSprint,
		IsDone:         todo.IsDone,
		Priority:       todo.Priority,
		Id:             todo.ID,
	}
}

func ListOfTodos(todos []Todo) []TodoGetPayload {
	todoList := []TodoGetPayload{}
	for _, todo := range todos {
		todoList = append(todoList, TodoGetPayload{
			Title:          todo.Title,
			SubTitle:       todo.SubTitle,
			EstimateSprint: todo.EstimateSprint,
			IsDone:         todo.IsDone,
			Priority:       todo.Priority,
			Id:             todo.ID,
		})
	}
	return todoList
}
