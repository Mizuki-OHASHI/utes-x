package usecase

import (
	"context"
	"utes-x-api/dao"
	"utes-x-api/model"
)

type User interface {
	GetMany(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, username string, email string) (*model.User, error)
}

type userUsecase struct {
	ud dao.User
}

func NewUserUsecase(ud dao.User) User {
	return &userUsecase{ud: ud}
}

func (u *userUsecase) GetMany(ctx context.Context) ([]model.User, error) {
	users, err := u.ud.GetMany(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) Create(ctx context.Context, username string, email string) (*model.User, error) {
	user := model.User{
		ID:       model.NewID(),
		Username: username,
		Email:    email,
	}
	newUser, err := u.ud.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
