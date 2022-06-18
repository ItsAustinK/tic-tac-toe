package database

import (
	"errors"
	gonanoid "github.com/matoous/go-nanoid"
)

type Game struct {
	Id          string
	Board       Board
	Players     []Player
	Actions     []Action
	CurPlayerId string
}

func NewGame(b Board) Game {
	nid, _ := gonanoid.Nanoid(16)
	return Game{
		Id:    nid,
		Board: b,
	}
}

func (g Game) IsValidPlayer(id string) bool {
	for i := range g.Players {
		if id == g.Players[i].Id {
			return true
		}
	}

	return false
}

func (g Game) IsPlayersTurn(id string) bool {
	return id == g.CurPlayerId
}

func (g Game) AddPlayerAction(a Action) error {
	if !g.IsPlayersTurn(a.PlayerId) {
		return errors.New("invalid action - not players turn")
	}

	g.Actions = append(g.Actions, a)
	return nil
}
