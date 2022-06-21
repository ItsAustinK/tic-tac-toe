package ticket

type Status string

const (
	Searching Status = "Searching"
	Complete  Status = "Completed"
)

type Ticket struct {
	Id            string
	UserId        string
	DateCreated   int64
	DateCompleted int64
	Status        string
	GameId        string
}
