package ticket

type Ticket struct {
	Id            string
	UserId        string
	DateCreated   int64
	DateCompleted int64
	Status        string
	GameId        string
}
