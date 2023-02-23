package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/sugimoto-ne/go_sample_app.git/clock"
	"github.com/sugimoto-ne/go_sample_app.git/config"
	"github.com/sugimoto-ne/go_sample_app.git/handler"
	"github.com/sugimoto-ne/go_sample_app.git/service"
	"github.com/sugimoto-ne/go_sample_app.git/store"
	"github.com/sugimoto-ne/go_sample_app.git/utils"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	// mux := http.NewServeMux()
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	v := validator.New()
	r := store.Repository{Clocker: clock.RealClocker{}}

	at := &handler.AddTask{
		Service: &service.AddTask{
			DB:   db,
			Repo: &r,
		},
		Validator: v,
	}
	mux.Post("/tasks", at.ServeHTTP)

	lt := &handler.ListTask{
		Service: &service.ListTask{
			DB:   db,
			Repo: &r,
		},
	}
	mux.Get("/tasks", lt.ServeHTTP)

	cv := utils.CustomValidator()
	ru := &handler.RegisterUser{
		Service: &service.RegisterUser{
			DB:   db,
			Repo: &r,
		},
		Validator: cv,
	}
	mux.Post("/register", ru.ServeHTTP)

	return mux, cleanup, nil
}
