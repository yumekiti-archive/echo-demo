package infrastructure

import (
	"github.com/yumekiti/echo-demo/domain/entity"
	"github.com/yumekiti/echo-demo/domain/repository"

	"gorm.io/gorm"
)

type taskRepository struct {
	Conn *gorm.DB
}

// NewTaskRepository task repositoryのコンストラクタ
func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &taskRepository{Conn: conn}
}

func (tr *taskRepository) GetAll() (*entity.Tasks, error) {
	tasks := &entity.Tasks{}

	if err := tr.Conn.Find(tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

// Create taskの保存
func (tr *taskRepository) Create(task *entity.Task) (*entity.Task, error) {
	if err := tr.Conn.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// FindByID taskをIDで取得
func (tr *taskRepository) FindByID(id int) (*entity.Task, error) {
	task := &entity.Task{ID: id}

	if err := tr.Conn.First(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// Update taskの更新
func (tr *taskRepository) Update(task *entity.Task) (*entity.Task, error) {
	if err := tr.Conn.Save(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// Delete taskの削除
func (tr *taskRepository) Delete(task *entity.Task) error {
	if err := tr.Conn.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
