package api

import (
	"P2/backend/user/application"
	"P2/backend/user/database"
	"context"
)

type User struct {
	Id               string
	Name             string
	AvatarUrl        string   // probably a s3 url
	ActiveUserIds    []string // could be its own db item
	CompletedUserIds []string // could be its own db item
}

func (u User) ToDbItem() database.User {
	item := database.User{
		Id:               u.Id,
		Name:             u.Name,
		AvatarUrl:        u.AvatarUrl,
		ActiveGameIds:    u.ActiveUserIds,
		CompletedGameIds: u.CompletedUserIds,
	}

	return item
}

func (u *User) FromDbItem(item database.User) {
	u.Id = item.Id
	u.Name = item.Name
	u.AvatarUrl = item.AvatarUrl
	u.ActiveUserIds = item.ActiveGameIds
	u.CompletedUserIds = item.CompletedGameIds
}

func GetUser(ctx context.Context, id string) (*User, error) {
	app := application.NewUserApp()
	u, err := app.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	dto := &User{}
	dto.FromDbItem(*u)
	return dto, nil
}

func Login(ctx context.Context, id string) (*User, error) {
	app := application.NewUserApp()
	u, err := app.Login(ctx, id)
	if err != nil {
		return nil, err
	}

	dto := &User{}
	dto.FromDbItem(*u)
	return dto, nil
}
