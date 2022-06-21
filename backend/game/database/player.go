package database

type Player struct {
	Id   string
	Name string
	Icon string
}

func NewPlayer(userId, userName string) Player {
	return Player{
		Id:   userId,
		Name: userName,
	}
}
