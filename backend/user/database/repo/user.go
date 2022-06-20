package repo

import (
	"P2/backend/user/database"
	"context"
)

type UserStorer interface {
	GetUser(ctx context.Context, id string) (*database.User, error)
	AddUser(ctx context.Context, user database.User) error
	UpdateUser(ctx context.Context, user database.User) error // preferable to update pieces of a user instead of whole obj
}

var udb *UserDatabase

// simple db
type UserDatabase struct {
	db map[string]database.User
}

func NewUserDatabase() *UserDatabase {
	if udb == nil {
		udb = &UserDatabase{db: map[string]database.User{}}
	}

	return udb
}

func (u *UserDatabase) GetUser(ctx context.Context, id string) (*database.User, error) {
	val, ok := u.db[id]
	if !ok {
		return nil, nil
	}
	return &val, nil
}

func (u *UserDatabase) AddUser(ctx context.Context, user database.User) error {
	u.db[user.Id] = user
	return nil
}

func (u *UserDatabase) UpdateUser(ctx context.Context, user database.User) error {
	u.db[user.Id] = user
	return nil
}
