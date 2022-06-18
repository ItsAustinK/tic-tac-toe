package api

import "P2/backend/game/database"

type Player struct {
	Id   string
	Name string
}

func (p Player) ToDbItem() database.Player {
	return database.Player{
		Id:   p.Id,
		Name: p.Name,
	}
}

func (p *Player) FromDbItem(item database.Player) {
	p.Name = item.Name
	p.Id = item.Id
}
