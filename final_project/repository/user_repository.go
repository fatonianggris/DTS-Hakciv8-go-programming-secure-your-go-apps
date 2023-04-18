package repository

import (
	"context"
	"database/sql"
	"go-programming-secure-your-go-apps/final_project/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User) (*domain.User, error)
	Login(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error)
	GetUserById(ctx context.Context, tx *sql.Tx, id string) (*domain.User, error)
	Update(ctx context.Context, user domain.User) (*domain.User, error)
	Delete(ctx context.Context, tx *sql.Tx, id string) error
}
