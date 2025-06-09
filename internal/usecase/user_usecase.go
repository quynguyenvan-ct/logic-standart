package usecase

import (
	"context"
	"golang/internal/entity"
	"golang/internal/repository"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
}

type UserUsecase struct {
	userrepo UserRepository
}

func NewUserUsecase() *UserUsecase {
    userRepo := repository.NewUserRepository()
    return &UserUsecase{
        userrepo: userRepo,
    }
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *entity.User) error {
	if err := uc.userrepo.Create(ctx,user); err != nil {
		return err
	}
	
	return nil
}

