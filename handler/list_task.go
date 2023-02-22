package handler

import (
	"net/http"

	"github.com/sugimoto-ne/go_sample_app.git/entity"
)

type ListTask struct {
	// Store *store.TaskStore
	// DB   *sqlx.DB
	// Repo *store.Repository
	Service ListTasksService
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := lt.Service.ListTasks(ctx)

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)

		return
	}

	rsp := []task{}

	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}

	RespondJSON(ctx, w, rsp, http.StatusOK)
}
