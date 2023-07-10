package service

import (
	todo "github.com/blazee5/todo-rest-api/models"
	"github.com/blazee5/todo-rest-api/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId int, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Update(userId int, listId int, input todo.UpdateListInput) error {
	return s.repo.Update(userId, listId, input)
}

func (s *TodoListService) Delete(userId int, listId int) error {
	return s.repo.Delete(userId, listId)
}