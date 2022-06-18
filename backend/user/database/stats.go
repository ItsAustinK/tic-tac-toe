package database

import gonanoid "github.com/matoous/go-nanoid"

type Stats struct {
	Id     string
	Wins   int
	Losses int
	Ties   int
}

func NewStats() Stats {
	nid, _ := gonanoid.Nanoid(16)
	return Stats{
		Id:     nid,
		Wins:   0,
		Losses: 0,
		Ties:   0,
	}
}
