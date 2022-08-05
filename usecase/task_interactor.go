package usecase

import "github.com/yumekiti/echo-demo/domain"

type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *TaskInteractor) TaskById(id int) (Task domain.Task, err error) {
	Task, err = interactor.TaskRepository.FindById(id)
	return
}

func (interactor *TaskInteractor) Tasks() (Tasks domain.Tasks, err error) {
	Tasks, err = interactor.TaskRepository.FindAll()
	return
}

func (interactor *TaskInteractor) Add(u domain.Task) (Task domain.Task, err error) {
	Task, err = interactor.TaskRepository.Store(u)
	return
}

func (interactor *TaskInteractor) Update(u domain.Task) (Task domain.Task, err error) {
	Task, err = interactor.TaskRepository.Update(u)
	return
}

func (interactor *TaskInteractor) DeleteById(u domain.Task) (err error) {
	err = interactor.TaskRepository.DeleteById(u)
	return
}
