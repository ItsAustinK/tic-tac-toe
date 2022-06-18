package api

import (
	"P2/backend/user/application"
	"P2/backend/user/database"
	"context"
)

type Stats struct {
	Id     string
	Wins   int
	Losses int
	Ties   int
}

func (s Stats) ToDbItem() database.Stats {
	item := database.Stats{
		Id:     s.Id,
		Wins:   s.Wins,
		Losses: s.Losses,
		Ties:   s.Ties,
	}

	return item
}

func (s *Stats) FromDbItem(item database.Stats) {
	s.Id = item.Id
	s.Wins = item.Wins
	s.Losses = item.Losses
	s.Ties = item.Ties
}

func GetStats(ctx context.Context, id string) (*Stats, error) {
	app := application.NewStatsApp()
	g, err := app.GetStats(ctx, id)
	if err != nil {
		return nil, err
	}

	dto := &Stats{}
	dto.FromDbItem(*g)
	return dto, nil
}
