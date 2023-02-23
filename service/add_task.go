package service

import (
	"context"
	"fmt"

	"github.com/sugimoto-ne/go_sample_app.git/entity"
	"github.com/sugimoto-ne/go_sample_app.git/store"
)

type AddTask struct {
	DB   store.Execer
	Repo TaskAdder
}

func (a *AddTask) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	t := &entity.Task{
		Title:  title,
		Status: entity.TaskStatusTodo,
	}

	err := a.Repo.AddTask(ctx, a.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	return t, nil
}
