package database

import (
	"errors"
	gonanoid "github.com/matoous/go-nanoid"
	"math/rand"
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
	Tie          Status = "tie"
)

type Game struct {
	Id           string
	Token        string
	CurPlayerIdx int
	WinnerIdx    int
	Presence     Presence
	Status       Status
	Board        Board
	Players      []Player
	Actions      []Action

	matchedUserIds []string // players that were matched
}

func NewGame(b Board, p Presence, uids []string) Game {
	nid, _ := gonanoid.Nanoid(16)
	t, _ := gonanoid.Nanoid(16)
	return Game{
		Id:           nid,
		Token:        t,
		CurPlayerIdx: 0,
		WinnerIdx:    -1,
		Presence:     p,
		Status:       Initializing,
		Board:        b,
		Players:      []Player{},
		Actions:      []Action{},

		matchedUserIds: uids,
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

func (g *Game) AddPlayer(p Player) {
	if g.Presence != Open && !g.isMatchedPlayer(p.Id) {
		return
	}

	if g.IsValidPlayer(p.Id) {
		return
	}

	var randChar rune
	for {
		randChar = 'A' + rune(rand.Intn(26))
		for i := range g.Players {
			if string(randChar) == g.Players[i].Icon {
				break
			}
		}
		break
	}

	p.Icon = string(randChar)
	g.Players = append(g.Players, p)

	if len(g.Players) == len(g.matchedUserIds) {
		g.Status = InProgress
	}
}

func (g *Game) AddPlayerAction(a Action) error {
	if !g.IsPlayersTurn(a.PlayerId) {
		return errors.New("invalid action - not players turn")
	}

	g.Actions = append(g.Actions, a)

	g.CurPlayerIdx = (g.CurPlayerIdx + 1) % len(g.Players)

	t, _ := gonanoid.Nanoid(16)
	g.Token = t

	return g.Board.AddAction(a)
}

func (g *Game) CheckForGameOver(a Action) {
	complete := g.Board.IsBoardComplete(a)
	if !complete {
		// check full board
		if len(g.Actions) == (g.Board.Row * g.Board.Col) {
			g.Status = Tie
		}

		return
	}

	g.WinnerIdx = g.getPlayerIdxById(a.PlayerId)
	g.Status = Complete
}

func (g Game) getPlayerIdxById(id string) int {
	for i := range g.Players {
		if id == g.Players[i].Id {
			return i
		}
	}

	return -1
}

func (g Game) isMatchedPlayer(id string) bool {
	for i := range g.matchedUserIds {
		if id == g.matchedUserIds[i] {
			return true
		}
	}

	return false
}
