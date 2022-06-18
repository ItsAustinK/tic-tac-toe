package database

import gonanoid "github.com/matoous/go-nanoid"

type User struct {
	Id               string
	Name             string
	AvatarUrl        string   // probably a s3 url
	ActiveGameIds    []string // could be its own db item
	CompletedGameIds []string // could be its own db item
}

func NewUser() User {
	nid, _ := gonanoid.Nanoid(16)
	return User{
		Id:               nid,
		Name:             "",
		AvatarUrl:        "",
		ActiveGameIds:    nil,
		CompletedGameIds: nil,
	}
}
