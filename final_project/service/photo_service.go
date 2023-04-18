package service

import (
	"context"
	"go-programming-secure-your-go-apps/final_project/model/domain"
)

type PhotoService interface {
	CreatePhoto(ctx context.Context, id string, photo domain.Photo) (*domain.Photo, error)
	GetPhotos() ([]*domain.Photo, error)
	GetPhotoById(ctx context.Context, id string) (*domain.Photo, error)
	UpdatePhoto(ctx context.Context, id string, photo domain.Photo) (*domain.Photo, error)
	DeletePhoto(ctx context.Context, id string) error
}
