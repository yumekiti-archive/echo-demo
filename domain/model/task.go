package model

import (
	"errors"

	"github.com/yumekiti/echo-demo/domain/entity"
)

type Task entity.Task

// NewTask taskのコンストラクタ
func NewTask(title, context string) (*Task, error) {
	if title == "" {
		return nil, errors.New("titleを入力してください")
	}

	task := &Task{
		Title:   title,
		Context: context,
	}

	return task, nil
}

// Set taskのセッター
func (t *Task) Set(title, context string) error {
	if title == "" {
		return errors.New("titleを入力してください")
	}

	t.Title = title
	t.Context = context

	return nil
}
