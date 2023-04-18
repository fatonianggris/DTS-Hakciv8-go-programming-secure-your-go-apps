package repository

import (
	"context"
	"database/sql"
	"go-programming-secure-your-go-apps/final_project/model/domain"
)

type PhotoRepository interface {
	CreatePhoto(ctx context.Context, tx *sql.Tx, id string, photo domain.Photo) (*domain.Photo, error)
	GetPhotos(tx *sql.Tx) ([]*domain.Photo, error)
	GetPhotoById(ctx context.Context, tx *sql.Tx, id string) (*domain.Photo, error)
	UpdatePhoto(ctx context.Context, tx *sql.Tx, id string, photo domain.Photo) (*domain.Photo, error)
	DeletePhoto(ctx context.Context, tx *sql.Tx, id string) error
}
