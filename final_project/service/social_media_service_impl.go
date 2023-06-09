package service

import (
	"context"
	"database/sql"
	"errors"
	"go-programming-secure-your-go-apps/final_project/exception"
	"go-programming-secure-your-go-apps/final_project/helper"
	"go-programming-secure-your-go-apps/final_project/model/domain"
	"go-programming-secure-your-go-apps/final_project/repository"
)

type SocialMediaServiceImpl struct {
	SocialMediaRepository repository.SocialMediaRepository
	DB                    *sql.DB
}

func NewSocialMediaService(socialMediaRepository repository.SocialMediaRepository, DB *sql.DB) SocialMediaService {
	return &SocialMediaServiceImpl{
		SocialMediaRepository: socialMediaRepository,
		DB:                    DB,
	}
}

func (smr *SocialMediaServiceImpl) CreateSocialMedia(ctx context.Context, id string, sm domain.SocialMediaInput) (*domain.SocialMedia, error) {
	if sm.Name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if sm.SocialMediaUrl == "" {
		return nil, errors.New("social media url cannot be empty")
	}

	var socialMedia domain.SocialMedia

	tx, err := smr.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	socialMedia.Name = sm.Name
	socialMedia.SocialMediaUrl = sm.SocialMediaUrl

	newSm, err := smr.SocialMediaRepository.CreateSocialMedia(ctx, tx, id, socialMedia)
	helper.PanicIfError(err)

	return newSm, nil
}

func (smr *SocialMediaServiceImpl) GetAllSocialMedia(ctx context.Context) ([]*domain.SocialMedia, error) {
	tx, err := smr.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getSmAll, err := smr.SocialMediaRepository.GetAllSocialMedia(ctx, tx)

	if err != nil {
		return nil, err
	}

	return getSmAll, nil
}

func (smr *SocialMediaServiceImpl) GetSocialMediaById(ctx context.Context, id string) (*domain.SocialMedia, error) {
	tx, err := smr.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getSmById, errGetById := smr.SocialMediaRepository.GetSocialMediaById(ctx, tx, id)
	if errGetById != nil {
		panic(exception.NewNotFoundError(errGetById.Error()))
	}

	return getSmById, nil
}

func (smr *SocialMediaServiceImpl) UpdateSocialMedia(ctx context.Context, id string, sm domain.SocialMediaInput) (*domain.SocialMedia, error) {
	if sm.Name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if sm.SocialMediaUrl == "" {
		return nil, errors.New("social media url cannot be empty")
	}

	var socialMedia domain.SocialMedia

	tx, err := smr.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	socialMedia.Name = sm.Name
	socialMedia.SocialMediaUrl = sm.SocialMediaUrl

	updateSm, err := smr.SocialMediaRepository.UpdateSocialMedia(ctx, tx, id, socialMedia)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return updateSm, nil
}

func (smr *SocialMediaServiceImpl) DeleteSocialMedia(ctx context.Context, id string) error {
	tx, err := smr.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	errDelete := smr.SocialMediaRepository.DeleteSocialMedia(ctx, tx, id)
	if errDelete != nil {
		panic(exception.NewNotFoundError(errDelete.Error()))
	}

	return nil
}
