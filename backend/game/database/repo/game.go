package repo

import (
	"P2/backend/game/database"
	"context"
)

type GameStorer interface {
	GetGame(ctx context.Context, id string) (*database.Game, error)
	AddGame(ctx context.Context, game database.Game) error
	UpdateGame(ctx context.Context, game database.Game) error // preferable to update pieces of a game instead of whole obj
}

var gdb *GameDatabase

// simple db
type GameDatabase struct {
	db map[string]database.Game
}

func NewGameDatabase() *GameDatabase {
	if gdb == nil {
		gdb = &GameDatabase{db: map[string]database.Game{}}
	}

	return gdb
}

func (g *GameDatabase) GetGame(ctx context.Context, id string) (*database.Game, error) {
	val, ok := g.db[id]
	if !ok {
		return nil, nil
	}
	return &val, nil
}

func (g *GameDatabase) AddGame(ctx context.Context, game database.Game) error {
	g.db[game.Id] = game
	return nil
}

func (g *GameDatabase) UpdateGame(ctx context.Context, game database.Game) error {
	g.db[game.Id] = game
	return nil
}
