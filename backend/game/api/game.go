package api

import (
	"P2/backend/game/application"
	"P2/backend/game/database"
	"context"
)

type Game struct {
	Id      string
	Board   Board
	Players []Player
	Actions []Action
}

func (g Game) ToDbItem() database.Game {
	item := database.Game{
		Id:    g.Id,
		Board: g.Board.ToDbItem(),
	}

	item.Players = make([]database.Player, len(g.Players))
	for i := range g.Players {
		item.Players[i] = g.Players[i].ToDbItem()
	}

	item.Actions = make([]database.Action, len(g.Actions))
	for i := range g.Actions {
		item.Actions[i] = g.Actions[i].ToDbItem()
	}

	return item
}

func (g *Game) FromDbItem(item database.Game) {
	g.Id = item.Id
	g.Board.FromDbItem(item.Board)

	g.Players = make([]Player, len(item.Players))
	for i := range item.Players {
		g.Players[i].FromDbItem(item.Players[i])
	}

	g.Actions = make([]Action, len(item.Actions))
	for i := range item.Actions {
		g.Actions[i].FromDbItem(item.Actions[i])
	}
}

func GetGame(ctx context.Context, id string) (*Game, error) {
	app := application.NewGameApp()
	g, err := app.GetGame(ctx, id)
	if err != nil {
		return nil, err
	}

	dto := &Game{}
	dto.FromDbItem(*g)
	return dto, nil
}

func CreateGame(ctx context.Context, board Board) (*Game, error) {
	item := board.ToDbItem()

	app := application.NewGameApp()
	g, err := app.CreateGame(ctx, item)
	if err != nil {
		return nil, err
	}

	dto := &Game{}
	dto.FromDbItem(*g)
	return dto, nil
}

func MakePlayerAction(ctx context.Context, id string, action Action) error {
	item := action.ToDbItem()

	app := application.NewGameApp()
	return app.MakePlayerAction(ctx, id, item)
}
