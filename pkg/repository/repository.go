package repository

import (
	todo "github.com/blazee5/todo-rest-api/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(item todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
	Update(userId int, listId int, input todo.UpdateListInput) error
	Delete(userId int, listId int) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}