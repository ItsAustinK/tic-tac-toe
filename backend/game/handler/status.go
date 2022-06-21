package handler

import (
	"P2/backend/game/api"
	"P2/backend/infrastructure/http"
	"errors"
	gohttp "net/http"
)

type StatusHandler struct{}

func (s StatusHandler) GetPath() string {
	return "/status"
}

func (s StatusHandler) ServeHTTP(w gohttp.ResponseWriter, r *gohttp.Request) {
	switch http.Method(r.Method) {
	case http.GET:
		s.getGameStatus(w, r)
	default:
		http.WriteError(w, 404, errors.New("unsupported method type"))
		return
	}
}

func (s StatusHandler) getGameStatus(w gohttp.ResponseWriter, r *gohttp.Request) {
	id, err := http.GetQueryParameter(r, "id")
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
		return
	}

	token, err := http.GetQueryParameter(r, "token")
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
		return
	}

	game, err := api.GetGameStatus(r.Context(), id, token)
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
		return
	}

	http.WriteResponse(w, gohttp.StatusOK, game)
}
