package api

import (
	"P2/backend/matchmaking/application"
	"context"
)

func QueueForMatch(ctx context.Context, id string) (*Ticket, error) {
	app := application.MatchmakerApp{}
	t, err := app.QueueForMatch(ctx, id)
	if err != nil {
		return nil, err
	}

	dto := &Ticket{}
	dto.FromDbItem(*t)
	return dto, nil
}
