package game

import (
	"fmt"
)

type Game struct {
	Id          string
	Token       string // updated every action
	CurPlayerId string
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
			fmt.Print(g.PlayersMap[g.Board.Pieces[i][j].PlayerId].Name)
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
