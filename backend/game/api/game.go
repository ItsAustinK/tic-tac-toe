package api

import (
	"P2/backend/game/application"
	"P2/backend/game/database"
	"context"
)

type Game struct {
	Id          string
	Token       string // updated every action
	CurPlayerId string
	WinnerId    string
	Presence    string
	Status      string
	Board       Board
	Players     []Player
	Actions     []Action
}

func (g Game) ToDbItem() database.Game {
	item := database.Game{
		Id:    g.Id,
		Board: g.Board.ToDbItem(),
	}

	item.Players = make([]database.Player, len(g.Players))
	for i := range g.Players {
		if g.CurPlayerId == g.Players[i].Id {
			item.CurPlayerIdx = i
		}

		if g.WinnerId == g.Players[i].Id {
			item.WinnerIdx = i
		}

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
	g.Token = item.Token
	g.CurPlayerId = item.Players[item.CurPlayerIdx].Id
	g.Presence = string(item.Presence)
	g.Status = string(item.Status)
	g.Board.FromDbItem(item.Board)

	if item.WinnerIdx != -1 {
		g.WinnerId = item.Players[item.WinnerIdx].Id
	}

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

func CreateGame(ctx context.Context, r, c, k int) (*Game, error) {
	app := application.NewGameApp()
	g, err := app.CreateGame(ctx, r, c, k, database.Invite, nil) // TODO: allow presence type to be provided & add creator as player
	if err != nil {
		return nil, err
	}

	dto := &Game{}
	dto.FromDbItem(*g)
	return dto, nil
}

func JoinGame(ctx context.Context, uid, gid string) (*Game, error) {
	app := application.NewGameApp()
	g, err := app.JoinGame(ctx, uid, gid)
	if err != nil {
		return nil, err
	}

	dto := &Game{}
	dto.FromDbItem(*g)
	return dto, nil
}

func MakePlayerAction(ctx context.Context, id, token string, action Action) (*Game, error) {
	item := action.ToDbItem()

	app := application.NewGameApp()
	g, err := app.MakePlayerAction(ctx, id, token, item)
	if err != nil {
		return nil, err
	}

	dto := &Game{}
	dto.FromDbItem(*g)
	return dto, nil
}

func GetGameStatus(ctx context.Context, id, token string) (*Game, error) {
	app := application.NewGameApp()
	g, err := app.GetGameStatus(ctx, id, token)
	if err != nil {
		return nil, err
	}

	dto := &Game{}
	dto.FromDbItem(*g)
	return dto, nil
}
