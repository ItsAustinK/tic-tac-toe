package application

import (
	"P2/backend/user/database"
	"P2/backend/user/database/repo"
	"context"
	"errors"
)

type UserApp struct {
	db repo.UserStorer
}

func NewUserApp() UserApp {
	return UserApp{
		db: repo.NewUserDatabase(),
	}
}

func (u UserApp) GetUser(ctx context.Context, id string) (*database.User, error) {
	user, err := u.db.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("failed to find database item")
	}

	return user, nil
}

func (u UserApp) AddUser(ctx context.Context) (*database.User, error) {
	user := database.NewUser()
	return &user, u.db.AddUser(ctx, user)
}

func (u UserApp) UpdateUser(ctx context.Context, user database.User) error {
	return u.db.UpdateUser(ctx, user)
}

func (u UserApp) Login(ctx context.Context, id string) (*database.User, error) {
	user, err := u.db.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	// create a new user
	return u.AddUser(ctx)
}
