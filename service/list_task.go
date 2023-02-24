package service

import (
	"context"
	"fmt"

	"github.com/sugimoto-ne/go_sample_app.git/auth"
	"github.com/sugimoto-ne/go_sample_app.git/entity"
	"github.com/sugimoto-ne/go_sample_app.git/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	ts, err := l.Repo.ListTasks(ctx, l.DB, id)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}
