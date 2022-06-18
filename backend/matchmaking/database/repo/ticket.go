package repo

import (
	"P2/backend/matchmaking/database"
	"context"
)

type TicketStorer interface {
	GetTicket(ctx context.Context, id string) (*database.Ticket, error)
	GetTickets(ctx context.Context, ids []string) ([]*database.Ticket, error)
	AddTicket(ctx context.Context, ticket database.Ticket) error
	UpdateTicket(ctx context.Context, ticket database.Ticket) error // preferable to update pieces of a ticket instead of whole obj
}

var tdb *TicketDatabase

// simple db
type TicketDatabase struct {
	db map[string]database.Ticket
}

func NewTicketDatabase() *TicketDatabase {
	if tdb == nil {
		tdb = &TicketDatabase{db: map[string]database.Ticket{}}
	}

	return tdb
}

func (t *TicketDatabase) GetTicket(ctx context.Context, id string) (*database.Ticket, error) {
	val := t.db[id]
	return &val, nil
}

func (t *TicketDatabase) GetTickets(ctx context.Context, ids []string) ([]*database.Ticket, error) {
	tickets := make([]*database.Ticket, len(ids))
	for i := range ids {
		val, ok := t.db[ids[i]]
		if !ok {
			continue
		}

		tickets[i] = &val
	}

	return tickets, nil
}

func (t *TicketDatabase) AddTicket(ctx context.Context, ticket database.Ticket) error {
	t.db[ticket.Id] = ticket
	return nil
}

func (t *TicketDatabase) UpdateTicket(ctx context.Context, ticket database.Ticket) error {
	t.db[ticket.Id] = ticket
	return nil
}

type TicketQueuer interface {
	QueueTicket(ctx context.Context, id string) error
	CheckTickets(ctx context.Context, matchSize int) ([]string, error)
}

var tq *TicketQueue

// simple queue
type TicketQueue struct {
	queue []string
}

func NewTicketQueue() *TicketQueue {
	if tq == nil {
		tq = &TicketQueue{queue: []string{}}
	}

	return tq
}

func (t *TicketQueue) QueueTicket(ctx context.Context, id string) error {
	// TODO: possible to queue more than once
	t.queue = append(t.queue, id)
	return nil
}

func (t *TicketQueue) CheckTickets(ctx context.Context, matchSize int) ([]string, error) {
	// found match
	if len(t.queue) >= matchSize {
		out := t.queue[0:matchSize]
		t.queue = t.queue[3:]

		return out, nil
	}

	return nil, nil
}
