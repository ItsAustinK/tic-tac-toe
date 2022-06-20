package database

import (
	"errors"
	gonanoid "github.com/matoous/go-nanoid"
)

type Game struct {
	Id           string
	Token        string
	CurPlayerIdx int
	WinnerIdx    int
	Board        Board
	Players      []Player
	Actions      []Action
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
	p := g.Players[g.CurPlayerIdx]
	return id == p.Id
}

func (g *Game) AddPlayerAction(a Action) error {
	if !g.IsPlayersTurn(a.PlayerId) {
		return errors.New("invalid action - not players turn")
	}

	g.Actions = append(g.Actions, a)

	g.CurPlayerIdx = (g.CurPlayerIdx + 1) % len(g.Players)

	return nil
}

func (g *Game) IsGameOver() bool {
	winnerId, complete := g.Board.IsBoardComplete()
	if !complete {
		return false
	}

	g.WinnerIdx = g.getPlayerIdxById(winnerId)
	return true
}

func (g Game) getPlayerIdxById(id string) int {
	for i := range g.Players {
		if id == g.Players[i].Id {
			return i
		}
	}

	return -1
}
