package service

import (
	"context"
	"go-programming-secure-your-go-apps/final_project/model/domain"
)

type SocialMediaService interface {
	CreateSocialMedia(ctx context.Context, id string, sm domain.SocialMediaInput) (*domain.SocialMedia, error)
	GetAllSocialMedia(ctx context.Context) ([]*domain.SocialMedia, error)
	GetSocialMediaById(ctx context.Context, id string) (*domain.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, id string, sm domain.SocialMediaInput) (*domain.SocialMedia, error)
	DeleteSocialMedia(ctx context.Context, id string) error
}
