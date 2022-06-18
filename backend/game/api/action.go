package api

import "P2/backend/game/database"

type Action struct {
	PlayerId string
	Position [2]int
}

func (a Action) ToDbItem() database.Action {
	return database.Action{
		PlayerId: a.PlayerId,
		Position: a.Position,
	}
}

func (a *Action) FromDbItem(item database.Action) {
	a.PlayerId = item.PlayerId
	a.Position = item.Position
}
