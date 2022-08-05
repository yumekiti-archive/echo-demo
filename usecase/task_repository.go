package usecase

import "github.com/yumekiti/echo-demo/domain"

type TaskRepository interface {
	FindById(id int) (domain.Task, error)
	FindAll() (domain.Tasks, error)
	Store(domain.Task) (domain.Task, error)
	Update(domain.Task) (domain.Task, error)
	DeleteById(domain.Task) error
}
