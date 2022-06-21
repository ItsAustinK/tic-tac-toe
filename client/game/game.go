package game

import (
	"fmt"
)

type Presence string

const (
	Open   Presence = "open"
	Closed Presence = "closed"
	Invite Presence = "invite"
)

type Status string

const (
	Initializing Status = "initializing"
	InProgress   Status = "in_progress"
	Complete     Status = "complete"
)

type Game struct {
	Id          string
	Token       string // updated every action
	CurPlayerId string
	WinnerId    string
	Presence    string
	Status      string
	Board       Board
	Players     []Player // this isn't necessary - could have better json unmarshalling to control what client models look like
	Actions     []Action

	PlayersMap map[string]Player // TODO: better json unmarshalling so this is populated initially and we don't need Players slice
}

func (g *Game) Init() {
	g.PlayersMap = map[string]Player{}
	for i := range g.Players {
		g.PlayersMap[g.Players[i].Id] = g.Players[i]
	}
}

func (g Game) Render() {
	for i := range g.Board.Pieces {
		for j := range g.Board.Pieces[i] {
			if g.PlayersMap[g.Board.Pieces[i][j].PlayerId].Id == "" { // empty piece
				fmt.Print(fmt.Sprintf("[%d]", i*len(g.Board.Pieces[i])+j))
			} else {
				fmt.Print(g.PlayersMap[g.Board.Pieces[i][j].PlayerId].Icon)
			}
		}
		fmt.Print("\n")
	}
}

func (g Game) GetPlayerById(id string) *Player {
	for i := range g.Players {
		if id == g.Players[i].Id {
			return &g.Players[i]
		}
	}

	return nil
}
