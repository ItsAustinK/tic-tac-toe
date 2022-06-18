package application

import (
	"P2/backend/matchmaking/database"
	"P2/backend/matchmaking/database/repo"
	"context"
)

type TicketApp struct {
	db repo.TicketStorer
}

func NewTicketApp() TicketApp {
	return TicketApp{
		db: repo.NewTicketDatabase(),
	}
}

func (t TicketApp) GetTicket(ctx context.Context, id string) (*database.Ticket, error) {
	return t.db.GetTicket(ctx, id)
}
