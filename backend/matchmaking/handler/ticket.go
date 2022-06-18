package handler

import (
	"P2/backend/infrastructure/http"
	"P2/backend/matchmaking/api"
	"errors"
	gohttp "net/http"
)

type TickerHandler struct{}

func (u TickerHandler) GetPath() string {
	return "/tickets"
}

func (u TickerHandler) ServeHTTP(w gohttp.ResponseWriter, r *gohttp.Request) {
	switch http.Method(r.Method) {
	case http.GET: // get a ticket
		id, err := http.GetQueryParameter(r, "id")
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		ticket, err := api.GetTicket(r.Context(), id)
		if err == nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
		}

		http.WriteResponse(w, gohttp.StatusOK, ticket)

	default:
		http.WriteError(w, 404, errors.New("unsupported method type"))
		return
	}
}
