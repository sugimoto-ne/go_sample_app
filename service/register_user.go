package service

import (
	"context"
	"fmt"

	"github.com/sugimoto-ne/go_sample_app.git/entity"
	"github.com/sugimoto-ne/go_sample_app.git/store"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	DB   store.Execer
	Repo UserRegisterer
}

func (ru *RegisterUser) RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &entity.User{
		Name:     name,
		Password: string(hashed),
		Role:     entity.Role(role),
	}

	err = ru.Repo.RegisterUser(ctx, ru.DB, u)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	return u, nil
}
