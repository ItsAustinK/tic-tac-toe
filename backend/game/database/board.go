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

// Ref: https://stackoverflow.com/questions/1056316/algorithm-for-determining-tic-tac-toe-game-over
func (b Board) IsBoardComplete(a Action) bool {
	//check col
	for i := 0; i < b.Col; i++ {
		if b.Pieces[a.Position[0]][i].PlayerId != a.PlayerId {
			break
		}

		if i == b.KVal-1 {
			return true
		}
	}

	//check row
	for i := 0; i < b.Row; i++ {
		if b.Pieces[i][a.Position[1]].PlayerId != a.PlayerId {
			break
		}

		if i == b.KVal-1 {
			return true
		}
	}

	//check diag
	if a.Position[0] == a.Position[1] {
		//we're on a diagonal
		for i, j := 0, 0; i < b.Row || j < b.Col; i, j = i+1, j+1 {
			if b.Pieces[i][j].PlayerId != a.PlayerId {
				break
			}
			if i == b.KVal-1 {
				return true
			}
		}
	}

	//check anti diag
	if a.Position[0]+a.Position[1] == b.KVal-1 {
		for i, j := 0, 0; i < b.Row && j < b.Col; i, j = i+1, j+1 {
			if b.Pieces[i][(b.KVal-1)-j].PlayerId != a.PlayerId {
				break
			}
			if i == b.KVal-1 {
				return true
			}
		}
	}

	return false
}
