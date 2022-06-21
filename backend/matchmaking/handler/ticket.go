package handler

import (
	"P2/backend/infrastructure/http"
	"P2/backend/matchmaking/api"
	"errors"
	gohttp "net/http"
)

type TicketHandler struct{}

func (u TicketHandler) GetPath() string {
	return "/tickets"
}

func (u TicketHandler) ServeHTTP(w gohttp.ResponseWriter, r *gohttp.Request) {
	switch http.Method(r.Method) {
	case http.GET:
		u.getTicket(w, r)
	default:
		http.WriteError(w, 404, errors.New("unsupported method type"))
		return
	}
}

func (u TicketHandler) getTicket(w gohttp.ResponseWriter, r *gohttp.Request) {
	id, err := http.GetQueryParameter(r, "id")
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
		return
	}

	ticket, err := api.GetTicket(r.Context(), id)
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
	}

	http.WriteResponse(w, gohttp.StatusOK, ticket)
}
