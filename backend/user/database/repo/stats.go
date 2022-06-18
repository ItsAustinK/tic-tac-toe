package repo

import (
	"P2/backend/user/database"
	"context"
)

type StatsStorer interface {
	GetStats(ctx context.Context, id string) (*database.Stats, error)
	AddStats(ctx context.Context, stats database.Stats) error
	UpdateStats(ctx context.Context, stats database.Stats) error // preferable to update pieces of a stats instead of whole obj
}

var sdb *StatsDatabase

// simple db
type StatsDatabase struct {
	db map[string]database.Stats
}

func NewStatsDatabase() *StatsDatabase {
	if sdb == nil {
		sdb = &StatsDatabase{db: map[string]database.Stats{}}
	}

	return sdb
}

func (s *StatsDatabase) GetStats(ctx context.Context, id string) (*database.Stats, error) {
	val := s.db[id]
	return &val, nil
}

func (s *StatsDatabase) AddStats(ctx context.Context, stats database.Stats) error {
	s.db[stats.Id] = stats
	return nil
}

func (s *StatsDatabase) UpdateStats(ctx context.Context, stats database.Stats) error {
	s.db[stats.Id] = stats
	return nil
}
