package game

type Board struct {
	Pieces [][]Action
	Row    int
	Col    int
	KVal   int // num items in a row to win
}
