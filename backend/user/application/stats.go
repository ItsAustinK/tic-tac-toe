package application

import (
	"P2/backend/user/database"
	"P2/backend/user/database/repo"
	"context"
	"errors"
)

type StatsApp struct {
	db repo.StatsStorer
}

func NewStatsApp() StatsApp {
	return StatsApp{
		db: repo.NewStatsDatabase(),
	}
}

func (s StatsApp) GetStats(ctx context.Context, id string) (*database.Stats, error) {
	stats, err := s.db.GetStats(ctx, id)
	if err != nil {
		return nil, err
	}

	if stats == nil {
		return nil, errors.New("failed to find database item")
	}

	return stats, nil
}

func (s StatsApp) CreateStats(ctx context.Context) (*database.Stats, error) {
	stats := database.NewStats()
	return &stats, s.db.AddStats(ctx, stats)
}

func (s StatsApp) UpdateStats(ctx context.Context, stats database.Stats) error {
	return s.db.UpdateStats(ctx, stats)
}
