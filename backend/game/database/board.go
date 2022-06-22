package database

import (
	"errors"
)

type Board struct {
	Pieces [][]Action
	Row    int
	Col    int
	KVal   int // num items in a row to win
}

func NewBoard(r, c, k int) Board {
	b := Board{
		Row:  r,
		Col:  c,
		KVal: k,
	}

	b.Pieces = make([][]Action, r)
	for i := range b.Pieces {
		b.Pieces[i] = make([]Action, c)
	}

	return b
}

func (b Board) IsPieceAvailable(position [2]int) bool {
	val := b.Pieces[position[0]][position[1]]
	return val.PlayerId == ""
}

func (b Board) AddAction(a Action) error {
	if !b.IsPieceAvailable(a.Position) {
		return errors.New("invalid action - board piece not available")
	}

	b.Pieces[a.Position[0]][a.Position[1]] = a
	return nil
}

func (b Board) IsBoardComplete() (string, bool) {
	isFull := true
	//winnerId := ""

	// TODO: check for winner
	for i := range b.Pieces {
		for j := range b.Pieces[i] {
			if b.IsPieceAvailable([2]int{i, j}) {
				isFull = false
			}
		}
	}

	if isFull {
		return "", true
	}

	//if winnerId != "" {
	//	return winnerId, true
	//}

	return "", false
}
