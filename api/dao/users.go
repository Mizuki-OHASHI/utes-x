package dao

import (
	"context"
	"database/sql"
	"utes-x-api/model"
	sqlboiler "utes-x-api/sqlboiler/entity"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/xerrors"
)

type User interface {
	GetMany(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, user model.User) (*model.User, error)
}

type useeDao struct {
	db *sql.DB
}

func NewUserDao(db *sql.DB) User {
	return &useeDao{db: db}
}

func (u *useeDao) GetMany(ctx context.Context) ([]model.User, error) {
	users, err := sqlboiler.Users().All(ctx, u.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to get users: %w", err)
	}
	return toUserModelSlice(users)
}

func (u *useeDao) Create(ctx context.Context, user model.User) (*model.User, error) {
	userDto := sqlboiler.User{
		ID:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}
	if err := userDto.Insert(ctx, u.db, boil.Infer()); err != nil {
		return nil, xerrors.Errorf("failed to create user: %w", err)
	}
	newUserDto, err := sqlboiler.Users(
		sqlboiler.UserWhere.ID.EQ(user.ID.String()),
	).One(ctx, u.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to fetch created user: %w", err)
	}
	return toUserModel(*newUserDto)
}
