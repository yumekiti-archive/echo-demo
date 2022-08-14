package repository

import (
	"github.com/yumekiti/echo-demo/domain/model"
)

type TaskRepository interface {
	GetAll() (*model.Tasks, error)
	Create(task *model.Task) (*model.Task, error)
	FindByID(id int) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(task *model.Task) error
}
