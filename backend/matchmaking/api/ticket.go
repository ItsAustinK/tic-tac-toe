package api

import (
	"P2/backend/matchmaking/application"
	"P2/backend/matchmaking/database"
	"context"
)

type Ticket struct {
	Id            string
	UserId        string
	DateCreated   int64
	DateCompleted int64
	Status        string
	GameId        string
}

func (t Ticket) ToDbItem() database.Ticket {
	item := database.Ticket{
		Id:            t.Id,
		UserId:        t.UserId,
		DateCreated:   t.DateCreated,
		DateCompleted: t.DateCompleted,
		Status:        database.Status(t.Status),
		GameId:        t.GameId,
	}

	return item
}

func (t *Ticket) FromDbItem(item database.Ticket) {
	t.Id = item.Id
	t.UserId = item.UserId
	t.DateCreated = item.DateCreated
	t.DateCompleted = item.DateCompleted
	t.Status = string(item.Status)
	t.GameId = item.GameId
}

func GetTicket(ctx context.Context, id string) (*Ticket, error) {
	app := application.NewTicketApp()
	t, err := app.GetTicket(ctx, id)
	if err != nil {
		return nil, err
	}

	dto := &Ticket{}
	dto.FromDbItem(*t)
	return dto, nil
}
