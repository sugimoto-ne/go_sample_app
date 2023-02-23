package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sugimoto-ne/go_sample_app.git/entity"
)

type AddTask struct {
	// Store     *store.TaskStore
	// DB        *sqlx.DB
	// Repo      *store.Repository
	Service   AddTaskService
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Title string `json:"title" validate:"required"`
	}

	// jsonのデコード
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)

		return
	}

	// バリデーションの検証
	err := at.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	// t := &entity.Task{
	// 	Title:  b.Title,
	// 	Status: entity.TaskStatusTodo,
	// }
	// err = at.Repo.AddTask(ctx, at.DB, t)
	t, err := at.Service.AddTask(ctx, b.Title)

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)

		return
	}

	rsp := struct {
		ID entity.TaskID `json:"id"`
	}{ID: t.ID}

	RespondJSON(ctx, w, rsp, http.StatusOK)

}
