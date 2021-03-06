package application

import (
	"P2/backend/game/database"
	"P2/backend/game/database/repo"
	"P2/backend/user/application"
	"context"
	"errors"
)

type GameApp struct {
	db      repo.GameStorer
	userApp application.UserApp
}

func NewGameApp() GameApp {
	return GameApp{
		db:      repo.NewGameDatabase(),
		userApp: application.NewUserApp(),
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

func (g GameApp) CreateGame(ctx context.Context, r, c, k int, p database.Presence, uids []string) (*database.Game, error) {
	board := database.NewBoard(r, c, k)
	game := database.NewGame(board, p, uids)
	return &game, g.db.AddGame(ctx, game)
}

func (g GameApp) JoinGame(ctx context.Context, uid, gid string) (*database.Game, error) {
	game, err := g.db.GetGame(ctx, gid)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, errors.New("failed to find game db item")
	}

	user, err := g.userApp.GetUser(ctx, uid)
	if err != nil {
		return nil, err

	}

	if user == nil {
		return nil, errors.New("failed to find user db item")
	}

	p := database.NewPlayer(user.Id, user.Name)
	game.AddPlayer(p)
	err = g.UpdateGame(ctx, *game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (g GameApp) UpdateGame(ctx context.Context, game database.Game) error {
	return g.db.UpdateGame(ctx, game)
}

func (g GameApp) MakePlayerAction(ctx context.Context, id, token string, action database.Action) (*database.Game, error) {
	game, err := g.GetGame(ctx, id)
	if err != nil {
		return nil, err
	}

	if token != game.Token {
		return nil, errors.New("invalid game token - player is out of sync with game")
	}

	if !game.IsValidPlayer(action.PlayerId) {
		return nil, errors.New("invalid action - player is not a part of the game")
	}

	if !game.IsPlayersTurn(action.PlayerId) {
		return nil, errors.New("invalid action - not player's turn")
	}

	if !game.Board.IsPieceAvailable(action.Position) {
		return nil, errors.New("invalid action - board piece not available")
	}

	err = game.AddPlayerAction(action)
	if err != nil {
		return nil, err
	}

	game.CheckForGameOver(action)

	return game, g.db.UpdateGame(ctx, *game)
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
