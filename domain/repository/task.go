package repository

import (
	"github.com/yumekiti/echo-demo/domain/entity"
)

type TaskRepository interface {
	GetAll() (*entity.Tasks, error)
	Create(task *entity.Task) (*entity.Task, error)
	FindByID(id int) (*entity.Task, error)
	Update(task *entity.Task) (*entity.Task, error)
	Delete(task *entity.Task) error
}
