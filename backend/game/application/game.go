package application

import (
	"P2/backend/game/database"
	"P2/backend/game/database/repo"
	"context"
	"errors"
)

type GameApp struct {
	db repo.GameStorer
}

func NewGameApp() GameApp {
	return GameApp{
		db: repo.NewGameDatabase(),
	}
}

func (g GameApp) GetGame(ctx context.Context, id string) (*database.Game, error) {
	game, err := g.db.GetGame(ctx, id)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, errors.New("failed to find database item")
	}

	return game, nil
}

func (g GameApp) CreateGame(ctx context.Context, board database.Board) (*database.Game, error) {
	game := database.NewGame(board)
	return &game, g.db.AddGame(ctx, game)
}

func (g GameApp) UpdateGame(ctx context.Context, game database.Game) error {
	return g.db.UpdateGame(ctx, game)
}

func (g GameApp) MakePlayerAction(ctx context.Context, id, token string, action database.Action) error {
	game, err := g.GetGame(ctx, id)
	if err != nil {
		return err
	}

	if token != game.Token {
		return errors.New("invalid game token - player is out of sync with game")
	}

	if !game.IsValidPlayer(action.PlayerId) {
		return errors.New("invalid action - player is not a part of the game")
	}

	if !game.IsPlayersTurn(action.PlayerId) {
		return errors.New("invalid action - not player's turn")
	}

	if !game.Board.IsPieceAvailable(action.Position) {
		return errors.New("invalid action - board piece not available")
	}

	err = game.AddPlayerAction(action)
	if err != nil {
		return err
	}

	err = game.Board.AddAction(action)
	if err != nil {
		return err
	}

	return g.db.UpdateGame(ctx, *game)
}

func (g GameApp) GetGameStatus(ctx context.Context, id, token string) (*database.Game, error) {
	game, err := g.GetGame(ctx, id)
	if err != nil {
		return nil, err
	}

	if token == game.Token {
		return nil, nil // no need to return updated game as they have latest
	}

	return game, nil
}
