package api

import "P2/backend/game/database"

type Board struct {
	KVal   int // num items in a row to win
	Pieces [][]Action
}

func (b Board) ToDbItem() database.Board {
	item := database.Board{
		KVal: b.KVal,
	}

	// TODO: look into optimizing
	item.Pieces = make([][]database.Action, len(b.Pieces))
	for i := range b.Pieces {
		item.Pieces[i] = make([]database.Action, len(b.Pieces[i]))
		for j := range b.Pieces[i] {
			item.Pieces[i][j] = b.Pieces[i][j].ToDbItem()
		}
	}

	return item
}

func (b *Board) FromDbItem(item database.Board) {
	b.KVal = item.KVal

	// TODO: look into optimizing
	b.Pieces = make([][]Action, len(item.Pieces))
	for i := range item.Pieces {
		b.Pieces[i] = make([]Action, len(item.Pieces[i]))
		for j := range item.Pieces[i] {
			a := &Action{}
			a.FromDbItem(item.Pieces[i][j])
			b.Pieces[i][j] = *a
		}
	}
}
