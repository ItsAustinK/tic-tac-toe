package user

type User struct {
	Id               string
	Name             string
	AvatarUrl        string   // probably a s3 url
	ActiveUserIds    []string // could be its own db item
	CompletedUserIds []string // could be its own db item
}
