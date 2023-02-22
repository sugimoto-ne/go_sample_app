package service

import (
	"context"

	"github.com/sugimoto-ne/go_sample_app.git/entity"
	"github.com/sugimoto-ne/go_sample_app.git/store"
)

type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}
