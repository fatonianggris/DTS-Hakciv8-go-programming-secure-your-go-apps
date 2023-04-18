package service

import (
	"context"
	"go-programming-secure-your-go-apps/final_project/model/domain"
)

type UserService interface {
	Register(ctx context.Context, user domain.User) (*domain.User, error)
	Login(ctx context.Context, login domain.UserLogin) (*domain.User, error)
	GetUserById(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, id string, user domain.UserLogin) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}
