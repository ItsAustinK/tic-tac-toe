package database

import (
	gonanoid "github.com/matoous/go-nanoid"
	"time"
)

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
	Status        Status
	GameId        string
}

func NewTicket(id string) Ticket {
	nid, _ := gonanoid.Nanoid(16)
	return Ticket{
		Id:          nid,
		UserId:      id,
		DateCreated: time.Now().UnixNano(),
		Status:      Searching,
	}
}

func (t *Ticket) SetCompletedStatus() {
	t.Status = Complete
	t.DateCompleted = time.Now().UnixNano()
}
