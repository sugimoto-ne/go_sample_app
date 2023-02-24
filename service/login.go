package service

import (
	"context"
	"fmt"

	"github.com/sugimoto-ne/go_sample_app.git/store"
)

type Login struct {
	Repo           UserGetter
	TokenGenerator TokenGenerator
	DB             store.Queryer
}

func (l *Login) Login(ctx context.Context, name, password string) (string, error) {

	u, err := l.Repo.GetUser(ctx, l.DB, name)
	if err != nil {
		return "", fmt.Errorf("failed to list: %w", err)
	}
	if err := u.ComparePassword(password); err != nil {
		return "", fmt.Errorf("wrong password: %w", err)
	}
	jwt, err := l.TokenGenerator.GenerateToken(ctx, *u)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	return string(jwt), nil
}
