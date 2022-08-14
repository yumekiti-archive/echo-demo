package model

import (
	"errors"

	"github.com/yumekiti/echo-demo/domain/entity"
)

// NewTask taskのコンストラクタ
func NewTask(title, context string) (*entity.Task, error) {
	if title == "" {
		return nil, errors.New("titleを入力してください")
	}

	task := &entity.Task{
		Title:   title,
		Context: context,
	}

	return task, nil
}

// Set taskのセッター
func Set(title, context string) error {
	if title == "" {
		return errors.New("titleを入力してください")
	}

	return nil
}
