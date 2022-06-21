package handler

import (
	"P2/backend/infrastructure/http"
	"P2/backend/matchmaking/api"
	"errors"
	"fmt"
	gohttp "net/http"
)

type MatchmakingHandler struct{}

func (m MatchmakingHandler) GetPath() string {
	return "/matchmake"
}

func (m MatchmakingHandler) ServeHTTP(w gohttp.ResponseWriter, r *gohttp.Request) {
	switch http.Method(r.Method) {
	case http.POST:
		m.queueForMatch(w, r)
	default:
		http.WriteError(w, 404, errors.New("unsupported method type"))
		return
	}
}

func (m MatchmakingHandler) queueForMatch(w gohttp.ResponseWriter, r *gohttp.Request) {
	id, err := http.GetQueryParameter(r, "id")
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
		return
	}
	fmt.Println(fmt.Sprintf("queuing user %s for match", id))

	ticket, err := api.QueueForMatch(r.Context(), id)
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
	}

	http.WriteResponse(w, gohttp.StatusOK, ticket)
}
