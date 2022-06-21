package application

import (
	"P2/backend/game/application"
	gameDb "P2/backend/game/database"
	"P2/backend/matchmaking/database"
	"P2/backend/matchmaking/database/repo"
	"context"
	"errors"
	"fmt"
	"log"
)

type MatchmakerApp struct {
	db      repo.TicketStorer
	queue   repo.TicketQueuer
	gameApp application.GameApp
}

func NewMatchmakerApp() MatchmakerApp {
	return MatchmakerApp{
		db:      repo.NewTicketDatabase(),
		queue:   repo.NewTicketQueue(),
		gameApp: application.NewGameApp(),
	}
}

func (m MatchmakerApp) QueueForMatch(ctx context.Context, id string) (*database.Ticket, error) {
	if id == "" {
		return nil, errors.New("invalid user id for queueing")
	}

	ticket := database.NewTicket(id)

	err := m.queue.QueueTicket(ctx, ticket.Id)
	if err != nil {
		return nil, err
	}

	err = m.db.AddTicket(ctx, ticket)
	if err != nil {
		return nil, err
	}

	// check to see if a match is made
	tickets, err := m.CheckForMatch(ctx)
	if err != nil {
		log.Print(fmt.Sprintf("something went wrong when checking for matches - err: %+v", err))
	}

	if tickets != nil {
		uids := make([]string, len(tickets))
		for i := range tickets {
			uids[i] = tickets[i].UserId
		}

		// create a game
		g, err := m.gameApp.CreateGame(ctx, 3, 3, 3, gameDb.Closed, uids)
		if err != nil {
			return nil, err
		}

		// update all tickets
		for i := range tickets {
			tickets[i].GameId = g.Id
			err = m.db.UpdateTicket(ctx, *tickets[i])
			if err != nil {
				return nil, err
			}

			// get the updated ticket to return
			if ticket.Id != tickets[i].Id {
				ticket = *tickets[i]
			}

		}
	}

	return &ticket, nil
}

func (m MatchmakerApp) CheckForMatch(ctx context.Context) ([]*database.Ticket, error) {
	ids, err := m.queue.CheckTickets(ctx, 2)
	if err != nil {
		return nil, err
	}

	// no match
	if ids == nil {
		return nil, err
	}

	tickets, err := m.db.GetTickets(ctx, ids)
	if err != nil {
		return nil, err
	}

	// some bad happened. we should have received all tickets
	if len(tickets) != len(ids) {
		log.Print(fmt.Sprintf("something went wrong. we're missing some tickets for the match"))
	}

	// update the tickets statuses
	for i := range tickets {
		tickets[i].SetCompletedStatus()
	}

	return tickets, nil
}
